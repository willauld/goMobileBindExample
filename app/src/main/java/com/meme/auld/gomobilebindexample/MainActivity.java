package com.meme.auld.gomobilebindexample;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.view.View;
import android.widget.Button;

import mygolib.JavaCallbackInterface;
import mygolib.Mygolib;

public class MainActivity extends AppCompatActivity {
    Button backtoJava;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Toolbar toolbar = findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);
    }

    public void handleClick(View v) throws Exception {
        Button but = (Button) v;
        String str = "Working...";
        but.setText(str);
        String resp = Mygolib.callgo("Calling go with me");
        but.setText(resp);
    }
    //*
    public class J implements JavaCallbackInterface {
        public void cjUpdate(String x) {
            // code you want called from Go.
            String str = "backtoJava "+x;
            backtoJava.setText(str);
        }
    }

    public void handleClick2(View v) throws Exception {
        backtoJava = (Button) v;
        String str = "Working2...";
        backtoJava.setText(str);
        String resp = Mygolib.callgocalljava(Mygolib.newUpdater(new J()));
        //backtoJava.setText(resp);
    }
    //*/
    /*
    public void handleClick2(View v) {
        backtoJava = (Button) v;
        String str = "Not implemented...";
        backtoJava.setText(str);
    }
    */
}
