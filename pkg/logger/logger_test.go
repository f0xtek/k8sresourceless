package logger

import (
	"testing"
)

type test struct {
	name     string
	expected string
}

type testData []NoResourceMsg

func TestFormat(t *testing.T) {
	t.Parallel()

	messages := testData{
		NoResourceMsg{CpuMissing: true, MemMissing: true, PodNamespace: "default", PodName: "test", PodOwner: "test", Content: ""},
		NoResourceMsg{CpuMissing: true, MemMissing: false, PodNamespace: "default", PodName: "test", PodOwner: "test", Content: ""},
		NoResourceMsg{CpuMissing: false, MemMissing: true, PodNamespace: "default", PodName: "test", PodOwner: "test", Content: ""},
	}

	tt := []test{
		{name: "cpu and memory missing", expected: "no cpu or memory resource requests defined"},
		{name: "cpu missing", expected: "no cpu resource requests defined"},
		{name: "memory missing", expected: "no memory resource requests defined"},
	}

	for i, message := range messages {
		message.format()
		if message.Content != tt[i].expected {
			t.Errorf("%v: got %v, want %v", tt[i].name, message.Content, tt[i].expected)
		}
	}
}

// func captureOutput(n *NoResourceMsg) string {
// 	buf := &bytes.Buffer{}
// 	log.SetOutput(buf)
// 	n.Log()
// 	log.SetOutput(os.Stderr)
// 	return buf.String()
// }

// TODO: FIX THIS TEST
// func TestLog(t *testing.T) {
// 	t.Parallel()

// 	messages := testData{
// 		NoResourceMsg{CpuMissing: true, MemMissing: true, PodNamespace: "default", PodName: "test", PodOwner: "test", Content: ""},
// 		NoResourceMsg{CpuMissing: true, MemMissing: false, PodNamespace: "default", PodName: "test", PodOwner: "test", Content: ""},
// 		NoResourceMsg{CpuMissing: false, MemMissing: true, PodNamespace: "default", PodName: "test", PodOwner: "test", Content: ""},
// 	}

// 	tt := []test{
// 		{name: "cpu and memory missing", expected: fmt.Sprintf("{\"level\":\"info\",\"namespace\":\"default\",\"name\":\"test\",\"owner\":\"test\",\"time\":%v,\"message\":\"no cpu or memory resource requests defined\"}", time.Now())},
// 		{name: "cpu missing", expected: fmt.Sprintf("{\"level\":\"info\",\"namespace\":\"default\",\"name\":\"test\",\"owner\":\"test\",\"time\":%v,\"message\":\"no cpu resource requests defined\"}", time.Now())},
// 		{name: "memory missing", expected: fmt.Sprintf("{\"level\":\"info\",\"namespace\":\"default\",\"name\":\"test\",\"owner\":\"test\",\"time\":%v,\"message\":\"no cpu resource requests defined\"}", time.Now())},
// 	}

// 	for i, message := range messages {
// 		output := captureOutput(&message) // how to properly capture output?
// 		if output != tt[i].expected {
// 			t.Errorf("%v: got %v, want %v", tt[i].name, output, tt[i].expected)
// 		}
// 	}
// }
