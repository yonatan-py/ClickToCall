package com.uruk.clicktocall

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import com.uruk.clicktocall.ui.theme.ClickToCallTheme

import android.util.Log
import com.google.firebase.firestore.ktx.firestore
import com.google.firebase.ktx.Firebase
import com.google.firebase.firestore.DocumentSnapshot
import com.google.android.gms.tasks.OnCompleteListener
import com.google.android.gms.tasks.Task

class MainActivity : ComponentActivity() {

    private val TAG = "MainActivity"
    val db = Firebase.firestore
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContent {
            ClickToCallTheme {
                // A surface container using the 'background' color from the theme
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    Greeting("Android")
                }
            }
        }
        db.collection("users")
            .document("UGcFEEx4kUzS7Y90NogY")
            .get()
            .addOnCompleteListener(OnCompleteListener { task: Task<DocumentSnapshot> ->
            if (task.isSuccessful) {
                val document: DocumentSnapshot = task.result!!

                if (document.exists()) {
                    // Document found, you can access the data using document.data
                    val data = document.data
                    // Do something with the data
                    Log.d(TAG, "DocumentSnapshot data: ${data}")
                } else {
                    // Document does not exist
                }
            } else {
                // An error occurred while fetching the document
                val exception = task.exception
                // Handle the error
            }
        })
    }
}

@Composable
fun Greeting(name: String, modifier: Modifier = Modifier) {
    Text(
        text = "Hello $name!",
        modifier = modifier
    )
}

@Preview(showBackground = true)
@Composable
fun GreetingPreview() {
    ClickToCallTheme {
        Greeting("Android")
    }
}