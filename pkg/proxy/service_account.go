package proxy

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/paralus/paralus/pkg/controller/apply"
	clientutil "github.com/paralus/paralus/pkg/controller/client"
	cruntime "github.com/paralus/paralus/pkg/controller/runtime"
	clusterv2 "github.com/paralus/paralus/proto/types/controller"
	"github.com/paralus/relay/pkg/relaylogger"
	"github.com/paralus/relay/pkg/utils"
	authenticationv1 "k8s.io/api/authentication/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

var (
	svclogger = relaylogger.NewLogger(utils.LogLevel).WithName("ServiceAccount")
	// ErrNotSAToken is returned when secret refered by service account is not of type service account token
	ErrNotSAToken = errors.New("secert is not of type ServiceAccountToken")
)

const (
	caCertKey = "ca.crt"
	tokenKey  = "token"
)

func getServiceAccountToken(ctx context.Context, name, namespace string) (string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", err
	}

	version, err := clientset.DiscoveryClient.ServerVersion()
	if err != nil {
		return "", err
	}

	major, err := strconv.ParseInt(strings.Trim(version.Major, "+_-"), 10, 64)
	if err != nil {
		return "", err
	}
	minor, err := strconv.ParseInt(strings.Trim(version.Minor, "+_-"), 10, 64)
	if err != nil {
		return "", err
	}

	if major >= 1 && minor >= 22 {
		tokenRequest, err := clientset.CoreV1().
			ServiceAccounts(namespace).
			CreateToken(ctx, name, &authenticationv1.TokenRequest{}, metav1.CreateOptions{})
		if err != nil {
			return "", err
		}

		return tokenRequest.Status.Token, nil
	}

	sa, err := clientset.CoreV1().
		ServiceAccounts(namespace).
		Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	for _, saSecret := range sa.Secrets {

		secret, err := clientset.CoreV1().Secrets(namespace).Get(ctx, saSecret.Name, metav1.GetOptions{})
		if err != nil {
			continue
		}

		if secret.Type == corev1.SecretTypeServiceAccountToken {
			return string(secret.Data[tokenKey]), nil
		}
	}

	return "", fmt.Errorf("service account %s/%s does not have secrets of type ServiceAccountToken", namespace, name)
}

// getServiceAccountSecret returns secret for the service account
func getServiceAccountSecret(ctx context.Context, c k8sclient.Client, name, namespace string) (*corev1.Secret, error) {
	var serviceAccount corev1.ServiceAccount
	var secret corev1.Secret

	err := c.Get(ctx, k8sclient.ObjectKey{
		Name:      name,
		Namespace: namespace,
	}, &serviceAccount)

	if err != nil {
		return nil, err
	}

	for _, saSecret := range serviceAccount.Secrets {
		err = c.Get(ctx, k8sclient.ObjectKey{
			Name:      saSecret.Name,
			Namespace: namespace,
		}, &secret)
		if err != nil {
			continue
		}

		if secret.Type == corev1.SecretTypeServiceAccountToken {
			return &secret, nil
		}
	}

	return nil, fmt.Errorf("service account %s/%s does not have secrets of type ServiceAccountToken", namespace, name)
}

//DeleteServiceAccount from cluster and cache
func DeleteServiceAccount(key, paralusAuthzSA, paralusAuthzRole, paralusAuthzRoleBind string, delCache bool) {
	svclogger.Debug(
		"DeleteServiceAccount",
		key,
		"not evicted", delCache,
	)
	dstrSA, err1 := base64.StdEncoding.DecodeString(paralusAuthzSA)
	dstrRole, err2 := base64.StdEncoding.DecodeString(paralusAuthzRole)
	dstrRB, err3 := base64.StdEncoding.DecodeString(paralusAuthzRoleBind)

	if err1 == nil && err2 == nil && err3 == nil {
		// delete sa,role,rolebinding
		svclogger.Debug(
			"delete service account",
			"yaml SA:",
			string(dstrSA),
			"yaml Role:",
			string(dstrRole),
			"yaml RoleBind:",
			string(dstrRB),
		)

		c, err := clientutil.New()
		if err != nil {
			svclogger.Error(
				err,
				"failed in clientutil new",
			)
			return
		}
		applier := apply.NewApplier(c)

		sa, err1 := getObject(dstrSA)
		role, err2 := getObject(dstrRole)
		rb, err3 := getObject(dstrRB)

		if err1 == nil && err2 == nil && err3 == nil {
			deleteAuthz(applier, rb)
			deleteAuthz(applier, role)
			deleteAuthz(applier, sa)
			if delCache {
				//delete from service cache
				utils.DeleteCache(utils.ServiceAccountCache, key)
			}

			return
		}

		svclogger.Error(
			nil,
			"failed to cleanup service account",
			"serviceaccount", string(dstrSA),
			"role", string(dstrRole),
			"rolebinding", string(dstrRB),
		)
	}
}

func getObject(yamlData []byte) (k8sclient.Object, error) {
	jb, err := yaml.YAMLToJSON(yamlData)
	if err != nil {
		return nil, err
	}

	var so clusterv2.StepObject
	err = json.Unmarshal(jb, &so)
	if err != nil {
		return nil, err
	}

	o, _, err := cruntime.ToUnstructuredObject(&so)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func deleteAuthz(applier apply.Applier, obj k8sclient.Object) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	err := applier.Delete(ctx, obj)
	if err != nil {
		svclogger.Error(
			err,
			"failed to delete sa/role/rolebinding",
		)
		return err
	}

	return nil
}
