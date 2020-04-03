package com.example.smsmanager;

import android.app.AlertDialog;
import android.app.PendingIntent;
import android.content.DialogInterface;
import android.content.Intent;
import android.database.Cursor;
import android.os.Bundle;
import android.provider.ContactsContract;
import android.support.v7.app.AppCompatActivity;
import android.telephony.SmsManager;
import android.view.View;
import android.view.ViewGroup;
import android.widget.BaseAdapter;
import android.widget.CheckBox;
import android.widget.EditText;
import android.widget.ListView;
import android.widget.Toast;

import java.util.ArrayList;
/**
*短信群发
*@author ZD
*created at 2017/8/7 16:35
*description：
*/ 
public class GroupSMSActivity extends AppCompatActivity {
    EditText numbers, content;

    SmsManager sManager;
    // 记录需要群发的号码列表
    ArrayList<String> sendList = new ArrayList<String>();

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_group_sms);
        sManager = SmsManager.getDefault();
        // 获取界面上的文本框、按钮组件
        numbers = (EditText) findViewById(R.id.numbers);
        content = (EditText) findViewById(R.id.content);
    }

    public void send(View view) {
        for (String number : sendList) {
            // 创建一个PendingIntent对象
            PendingIntent pi = PendingIntent.getActivity(
                    GroupSMSActivity.this, 0, new Intent(), 0);
            // 发送短信
            sManager.sendTextMessage(number, null, content
                    .getText().toString(), pi, null);
        }
        // 提示短信群发完成
        Toast.makeText(GroupSMSActivity.this, "短信群发完成"
                , Toast.LENGTH_SHORT).show();
    }


    public void select(View view) {
// 查询联系人的电话号码
        final Cursor cursor = getContentResolver().query(
                ContactsContract.CommonDataKinds.Phone.CONTENT_URI,
                null, null, null, null);
        BaseAdapter adapter = new BaseAdapter() {
            @Override
            public int getCount() {
                return cursor.getCount();
            }

            @Override
            public Object getItem(int position) {
                return position;
            }

            @Override
            public long getItemId(int position) {
                return position;
            }

            @Override
            public View getView(int position, View convertView,
                                ViewGroup parent) {
                cursor.moveToPosition(position);
                CheckBox rb = new CheckBox(GroupSMSActivity.this);
                // 获取联系人的电话号码，并去掉中间的中画线、空格
                String number = cursor
                        .getString(cursor.getColumnIndex(ContactsContract
                                .CommonDataKinds.Phone.NUMBER))
                        .replace("-", "")
                        .replace(" ", "");
                rb.setText(number);
                // 如果该号码已经被加入发送人名单，默认勾选该号码
                if (isChecked(number)) {
                    rb.setChecked(true);
                }
                return rb;
            }
        };
        // 加载list.xml布局文件对应的View
        View selectView = getLayoutInflater().inflate(
                R.layout.list, null);
        // 获取selectView中的名为list的ListView组件
        final ListView listView = (ListView) selectView
                .findViewById(R.id.list);
        listView.setAdapter(adapter);
        new AlertDialog.Builder(GroupSMSActivity.this)
                .setView(selectView)
                .setPositiveButton("确定",
                        new DialogInterface.OnClickListener() {
                            @Override
                            public void onClick(DialogInterface dialog,
                                                int which) {
                                // 清空sendList集合
                                sendList.clear();
                                // 遍历listView组件的每个列表项
                                for (int i = 0; i < listView.getCount(); i++) {
                                    CheckBox checkBox = (CheckBox) listView
                                            .getChildAt(i);
                                    // 如果该列表项被勾选
                                    if (checkBox.isChecked()) {
                                        // 添加该列表项的电话号码
                                        sendList.add(checkBox.getText()
                                                .toString());
                                    }
                                }
                                numbers.setText(sendList.toString());
                            }
                        }).show();

    }

    // 判断某个电话号码是否已在群发范围内
    public boolean isChecked(String phone) {
        for (String s1 : sendList) {
            if (s1.equals(phone)) {
                return true;
            }
        }
        return false;
    }
}