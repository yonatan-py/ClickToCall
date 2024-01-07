package com.uruk.clicktocall

import android.Manifest
import android.content.pm.PackageManager
import android.os.AsyncTask
import android.os.Build
import android.os.Bundle
import android.util.Log
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Modifier
import androidx.compose.ui.tooling.preview.Preview
import androidx.core.content.ContextCompat
import com.google.android.gms.tasks.OnCompleteListener
import com.google.firebase.messaging.FirebaseMessaging
import com.uruk.clicktocall.ui.theme.ClickToCallTheme
import kotlinx.coroutines.launch
import okhttp3.MediaType.Companion.toMediaTypeOrNull
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.Response
import org.json.JSONObject
import java.io.IOException


class MainActivity : ComponentActivity() {

    private val TAG = "MainActivity"
    private var code = ""
    private var token = ""

    private val requestPermissionLauncher = registerForActivityResult(
        ActivityResultContracts.RequestPermission(),
    ) { isGranted: Boolean ->
        if (isGranted) {
            // FCM SDK (and your app) can post notifications.
        } else {
            // TODO: Inform user that that your app will not show notifications.
        }
    }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        code = "123"
        setContent {
            ClickToCallTheme {
                // A surface container using the 'background' color from the theme
                Surface(
                    modifier = Modifier.fillMaxSize(),
                    color = MaterialTheme.colorScheme.background
                ) {
                    val coroutineScope = rememberCoroutineScope()
                    LaunchedEffect(code, token) {
                        Log.i(TAG, "code: $code")
                        Log.i(TAG, "token: $token")
                        coroutineScope.launch {
                            SendCodeToServer().execute(token, code)
                        }
                    }
                    ShowCode(code)
                }
            }
        }

        FirebaseMessaging.getInstance().token.addOnCompleteListener(OnCompleteListener { task ->
            if (!task.isSuccessful) {
                Log.w(TAG, "Fetching FCM registration token failed", task.exception)
                return@OnCompleteListener
            }
            token = task.result
            val msg = String.format("token: %s", token)
            Log.d(TAG, msg)
        })
        askNotificationPermission()
    }
    private fun askNotificationPermission() {
        // This is only necessary for API level >= 33 (TIRAMISU)
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU) {
            if (ContextCompat.checkSelfPermission(this, Manifest.permission.POST_NOTIFICATIONS) ==
                PackageManager.PERMISSION_GRANTED
            ) {
                // FCM SDK (and your app) can post notifications.
            } else if (shouldShowRequestPermissionRationale(Manifest.permission.POST_NOTIFICATIONS)) {
                // TODO: display an educational UI explaining to the user the features that will be enabled
                //       by them granting the POST_NOTIFICATION permission. This UI should provide the user
                //       "OK" and "No thanks" buttons. If the user selects "OK," directly request the permission.
                //       If the user selects "No thanks," allow the user to continue without notifications.
            } else {
                // Directly ask for the permission
                requestPermissionLauncher.launch(Manifest.permission.POST_NOTIFICATIONS)
            }
        }
    }
}

class SendCodeToServer : AsyncTask<String, Void, Response?>() {

    @Deprecated("Deprecated in Java")
    override fun doInBackground(vararg params: String): Response? {
        try {
            val baseUrl = BuildConfig.API_URL
            val url = "${baseUrl}/code"
            Log.d("url", url)
            val codePayload = JSONObject()
            val token = params[0]
            Log.d("token", token)
            val code = params[1]
            codePayload.put("code", code)
            codePayload.put("androidToken", token)
            val jsonString = codePayload.toString()
            val client = OkHttpClient()

            val request = Request.Builder()
                .url(url)
                .post(jsonString.toRequestBody("application/json; charset=utf-8".toMediaTypeOrNull()))
                .build()

            client.newCall(request).execute().use { response ->
                Log.d("response", response.toString())
                if (!response.isSuccessful) {
                    throw IOException("Unexpected code $response")
                }

                for ((name, value) in response.headers) {
                    println("$name: $value")
                }

                println(response.body!!.string())
            }
            return null
        } catch (e: Exception) {
            Log.e("error", e.toString())
            return null
        }
    }

    @Deprecated("Deprecated in Java")
    override fun onPostExecute(result: Response?) {
        // Update UI or perform any post-execution tasks here
        // This method runs on the main thread
    }
}

@Composable
fun ShowCode(code: String) {
    Column {
        Text(code)
    }
}

@Preview(showBackground = true)
@Composable
fun DefaultPreview() {
    ClickToCallTheme {
        ShowCode(code = "123")
    }
}