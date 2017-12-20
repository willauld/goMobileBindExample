# goMobileBindExample

mygolibcode // contains test code for mygolib
    mygolib // contains the example go lib code 
    
To build and test the library all in go: 
    go install
    go run try2.go

To build the android library:
    gomobile bind -target=android
    copy the resulting mygolib.aar to android lib directory

    In your project gradle file in the allprojects.repositories tag, add the following :

        flatDir {
            dirs 'libs'
        }
        
    In your app module, make sure you have a libs folder with the aar file in it. Then add the compile line in the dependencies tag of your app module gradle file:

        dependencies {
            compile 'package:name:version@aar'
        }
        I use: compile 'go.mygolib.gojni:mygolib:@aar'

Use android studio to build and test the mobile app.

Current status:
    App works with Mygolib.callgo() :: first button
    Last commit seems have the call from go to java working

        Was failing with to build with Mygolib.callgocalljava()
