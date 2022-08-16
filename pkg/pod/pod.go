package pod

import (
	"fmt"

	"github.com/f0xtek/k8sresourcess/pkg/logger"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type PodLoggingController struct {
	informerFactory informers.SharedInformerFactory
	podInformer     coreinformers.PodInformer
}

func (c *PodLoggingController) Run(stopCh chan struct{}) error {
	c.informerFactory.Start(stopCh)
	if !cache.WaitForCacheSync(stopCh, c.podInformer.Informer().HasSynced) {
		return fmt.Errorf("failed to sync")
	}
	return nil
}

func getOwner(p *v1.Pod) string {
	var owner string
	v, found := p.Labels["owner"]
	if found {
		owner = v
	} else {
		owner = "no owner"
	}
	return owner
}

func (c *PodLoggingController) podAdd(obj interface{}) {
	pod := obj.(*v1.Pod)
	podOwner := getOwner(pod)

	for _, container := range pod.Spec.Containers {
		var cpuRequestMissing, memoryRequestMissing bool
		if container.Resources.Requests.Cpu().IsZero() {
			cpuRequestMissing = true
		}
		if container.Resources.Requests.Memory().IsZero() {
			memoryRequestMissing = true
		}
		msg := logger.NoResourceMsg{
			CpuMissing:   cpuRequestMissing,
			MemMissing:   memoryRequestMissing,
			PodNamespace: pod.Namespace,
			PodName:      pod.Name,
			PodOwner:     podOwner,
			Content:      "",
		}
		msg.Log()
	}
}

func NewPodLoggingController(informerFactory informers.SharedInformerFactory) *PodLoggingController {
	podInformer := informerFactory.Core().V1().Pods()

	c := &PodLoggingController{
		informerFactory: informerFactory,
		podInformer:     podInformer,
	}
	podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: c.podAdd,
		},
	)
	return c
}
