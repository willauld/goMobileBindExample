package mygolib

import "fmt"

/*
Two main examples on how the callback into java should be done. First:
https://medium.com/@matryer/tutorial-calling-go-code-from-swift-on-ios-and-vice-versa-with-gomobile-7925620c17a4

Second Example explenation given in https://groups.google.com/forum/#!topic/Golang-Nuts/DRpXR6P-eG8

"That said, I think your project is very interesting and I hope you
keep working on it. One thing I noticed on a quick look through your
code: you may not realize that gobind will manage callbacks from Java
into Go. That is if your Go package defines: "

type Callback interface {
  CallJava(x T) U
}

func Register(c Callback) {
  c.CallJava(x) // c can be stored for later use
}
Then in Java you can subclass the generated class and implement the
CallJava method:

public class J extends gopkg.Callback {
    public CallJava(T x) U {
        // code you want called from Go.
    }
}
...
gopkg.Register(new J());
*/

type JavaCallbackInterface interface {
	CJUpdate(x string)
}

type Updater struct {
	javaCall JavaCallbackInterface
}

var iterationProgress *Updater

func NewUpdater(c JavaCallbackInterface) *Updater {
	return &Updater{javaCall: c}
}

func (u Updater) Update(x string) {
	u.javaCall.CJUpdate(x)
}

func CallbackToJava() {
	iterationProgress.Update("String past to Java method from go")
}

func Callgo(s string) (string, error) {
	str := fmt.Sprintf("String from calling a go function from java with string: \"%s\"\n", s)
	return str, nil
}

func Callgocalljava(f *Updater) (string, error) {
	if f != nil {
		iterationProgress = f //f = NewUpdater(y)
	}
	CallbackToJava()
	str := "String returned from calling a go function from java"
	return str, nil
}
