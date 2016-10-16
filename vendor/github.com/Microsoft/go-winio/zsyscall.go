// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package winio

import "unsafe"
import "syscall"

var _ unsafe.Pointer

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

	procCancelIoEx                                           = modkernel32.NewProc("CancelIoEx")
	procCreateIoCompletionPort                               = modkernel32.NewProc("CreateIoCompletionPort")
	procGetQueuedCompletionStatus                            = modkernel32.NewProc("GetQueuedCompletionStatus")
	procSetFileCompletionNotificationModes                   = modkernel32.NewProc("SetFileCompletionNotificationModes")
	procConnectNamedPipe                                     = modkernel32.NewProc("ConnectNamedPipe")
	procCreateNamedPipeW                                     = modkernel32.NewProc("CreateNamedPipeW")
	procCreateFileW                                          = modkernel32.NewProc("CreateFileW")
	procWaitNamedPipeW                                       = modkernel32.NewProc("WaitNamedPipeW")
	procLookupAccountNameW                                   = modadvapi32.NewProc("LookupAccountNameW")
	procConvertSidToStringSidW                               = modadvapi32.NewProc("ConvertSidToStringSidW")
	procConvertStringSecurityDescriptorToSecurityDescriptorW = modadvapi32.NewProc("ConvertStringSecurityDescriptorToSecurityDescriptorW")
	procConvertSecurityDescriptorToStringSecurityDescriptorW = modadvapi32.NewProc("ConvertSecurityDescriptorToStringSecurityDescriptorW")
	procLocalFree                                            = modkernel32.NewProc("LocalFree")
	procGetSecurityDescriptorLength                          = modadvapi32.NewProc("GetSecurityDescriptorLength")
	procGetFileInformationByHandleEx                         = modkernel32.NewProc("GetFileInformationByHandleEx")
	procSetFileInformationByHandle                           = modkernel32.NewProc("SetFileInformationByHandle")
	procAdjustTokenPrivileges                                = modadvapi32.NewProc("AdjustTokenPrivileges")
	procImpersonateSelf                                      = modadvapi32.NewProc("ImpersonateSelf")
	procRevertToSelf                                         = modadvapi32.NewProc("RevertToSelf")
	procOpenThreadToken                                      = modadvapi32.NewProc("OpenThreadToken")
	procGetCurrentThread                                     = modkernel32.NewProc("GetCurrentThread")
	procLookupPrivilegeValueW                                = modadvapi32.NewProc("LookupPrivilegeValueW")
	procLookupPrivilegeNameW                                 = modadvapi32.NewProc("LookupPrivilegeNameW")
	procLookupPrivilegeDisplayNameW                          = modadvapi32.NewProc("LookupPrivilegeDisplayNameW")
	procBackupRead                                           = modkernel32.NewProc("BackupRead")
	procBackupWrite                                          = modkernel32.NewProc("BackupWrite")
)

func cancelIoEx(file syscall.Handle, o *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procCancelIoEx.Addr(), 2, uintptr(file), uintptr(unsafe.Pointer(o)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func createIoCompletionPort(file syscall.Handle, port syscall.Handle, key uintptr, threadCount uint32) (newport syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateIoCompletionPort.Addr(), 4, uintptr(file), uintptr(port), uintptr(key), uintptr(threadCount), 0, 0)
	newport = syscall.Handle(r0)
	if newport == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getQueuedCompletionStatus(port syscall.Handle, bytes *uint32, key *uintptr, o **ioOperation, timeout uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetQueuedCompletionStatus.Addr(), 5, uintptr(port), uintptr(unsafe.Pointer(bytes)), uintptr(unsafe.Pointer(key)), uintptr(unsafe.Pointer(o)), uintptr(timeout), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setFileCompletionNotificationModes(h syscall.Handle, flags uint8) (err error) {
	r1, _, e1 := syscall.Syscall(procSetFileCompletionNotificationModes.Addr(), 2, uintptr(h), uintptr(flags), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func connectNamedPipe(pipe syscall.Handle, o *syscall.Overlapped) (err error) {
	r1, _, e1 := syscall.Syscall(procConnectNamedPipe.Addr(), 2, uintptr(pipe), uintptr(unsafe.Pointer(o)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func createNamedPipe(name string, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *securityAttributes) (handle syscall.Handle, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	return _createNamedPipe(_p0, flags, pipeMode, maxInstances, outSize, inSize, defaultTimeout, sa)
}

func _createNamedPipe(name *uint16, flags uint32, pipeMode uint32, maxInstances uint32, outSize uint32, inSize uint32, defaultTimeout uint32, sa *securityAttributes) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateNamedPipeW.Addr(), 8, uintptr(unsafe.Pointer(name)), uintptr(flags), uintptr(pipeMode), uintptr(maxInstances), uintptr(outSize), uintptr(inSize), uintptr(defaultTimeout), uintptr(unsafe.Pointer(sa)), 0)
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func createFile(name string, access uint32, mode uint32, sa *securityAttributes, createmode uint32, attrs uint32, templatefile syscall.Handle) (handle syscall.Handle, err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	return _createFile(_p0, access, mode, sa, createmode, attrs, templatefile)
}

func _createFile(name *uint16, access uint32, mode uint32, sa *securityAttributes, createmode uint32, attrs uint32, templatefile syscall.Handle) (handle syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall9(procCreateFileW.Addr(), 7, uintptr(unsafe.Pointer(name)), uintptr(access), uintptr(mode), uintptr(unsafe.Pointer(sa)), uintptr(createmode), uintptr(attrs), uintptr(templatefile), 0, 0)
	handle = syscall.Handle(r0)
	if handle == syscall.InvalidHandle {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func waitNamedPipe(name string, timeout uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	return _waitNamedPipe(_p0, timeout)
}

func _waitNamedPipe(name *uint16, timeout uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procWaitNamedPipeW.Addr(), 2, uintptr(unsafe.Pointer(name)), uintptr(timeout), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func lookupAccountName(systemName *uint16, accountName string, sid *byte, sidSize *uint32, refDomain *uint16, refDomainSize *uint32, sidNameUse *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(accountName)
	if err != nil {
		return
	}
	return _lookupAccountName(systemName, _p0, sid, sidSize, refDomain, refDomainSize, sidNameUse)
}

func _lookupAccountName(systemName *uint16, accountName *uint16, sid *byte, sidSize *uint32, refDomain *uint16, refDomainSize *uint32, sidNameUse *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procLookupAccountNameW.Addr(), 7, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(accountName)), uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(sidSize)), uintptr(unsafe.Pointer(refDomain)), uintptr(unsafe.Pointer(refDomainSize)), uintptr(unsafe.Pointer(sidNameUse)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func convertSidToStringSid(sid *byte, str **uint16) (err error) {
	r1, _, e1 := syscall.Syscall(procConvertSidToStringSidW.Addr(), 2, uintptr(unsafe.Pointer(sid)), uintptr(unsafe.Pointer(str)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func convertStringSecurityDescriptorToSecurityDescriptor(str string, revision uint32, sd *uintptr, size *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(str)
	if err != nil {
		return
	}
	return _convertStringSecurityDescriptorToSecurityDescriptor(_p0, revision, sd, size)
}

func _convertStringSecurityDescriptorToSecurityDescriptor(str *uint16, revision uint32, sd *uintptr, size *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procConvertStringSecurityDescriptorToSecurityDescriptorW.Addr(), 4, uintptr(unsafe.Pointer(str)), uintptr(revision), uintptr(unsafe.Pointer(sd)), uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func convertSecurityDescriptorToStringSecurityDescriptor(sd *byte, revision uint32, secInfo uint32, sddl **uint16, sddlSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procConvertSecurityDescriptorToStringSecurityDescriptorW.Addr(), 5, uintptr(unsafe.Pointer(sd)), uintptr(revision), uintptr(secInfo), uintptr(unsafe.Pointer(sddl)), uintptr(unsafe.Pointer(sddlSize)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func localFree(mem uintptr) {
	syscall.Syscall(procLocalFree.Addr(), 1, uintptr(mem), 0, 0)
	return
}

func getSecurityDescriptorLength(sd uintptr) (len uint32) {
	r0, _, _ := syscall.Syscall(procGetSecurityDescriptorLength.Addr(), 1, uintptr(sd), 0, 0)
	len = uint32(r0)
	return
}

func getFileInformationByHandleEx(h syscall.Handle, class uint32, buffer *byte, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procGetFileInformationByHandleEx.Addr(), 4, uintptr(h), uintptr(class), uintptr(unsafe.Pointer(buffer)), uintptr(size), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setFileInformationByHandle(h syscall.Handle, class uint32, buffer *byte, size uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetFileInformationByHandle.Addr(), 4, uintptr(h), uintptr(class), uintptr(unsafe.Pointer(buffer)), uintptr(size), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func adjustTokenPrivileges(token syscall.Handle, releaseAll bool, input *byte, outputSize uint32, output *byte, requiredSize *uint32) (success bool, err error) {
	var _p0 uint32
	if releaseAll {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, e1 := syscall.Syscall6(procAdjustTokenPrivileges.Addr(), 6, uintptr(token), uintptr(_p0), uintptr(unsafe.Pointer(input)), uintptr(outputSize), uintptr(unsafe.Pointer(output)), uintptr(unsafe.Pointer(requiredSize)))
	success = r0 != 0
	if true {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func impersonateSelf(level uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procImpersonateSelf.Addr(), 1, uintptr(level), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func revertToSelf() (err error) {
	r1, _, e1 := syscall.Syscall(procRevertToSelf.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func openThreadToken(thread syscall.Handle, accessMask uint32, openAsSelf bool, token *syscall.Handle) (err error) {
	var _p0 uint32
	if openAsSelf {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall6(procOpenThreadToken.Addr(), 4, uintptr(thread), uintptr(accessMask), uintptr(_p0), uintptr(unsafe.Pointer(token)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func getCurrentThread() (h syscall.Handle) {
	r0, _, _ := syscall.Syscall(procGetCurrentThread.Addr(), 0, 0, 0, 0)
	h = syscall.Handle(r0)
	return
}

func lookupPrivilegeValue(systemName string, name string, luid *uint64) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return
	}
	var _p1 *uint16
	_p1, err = syscall.UTF16PtrFromString(name)
	if err != nil {
		return
	}
	return _lookupPrivilegeValue(_p0, _p1, luid)
}

func _lookupPrivilegeValue(systemName *uint16, name *uint16, luid *uint64) (err error) {
	r1, _, e1 := syscall.Syscall(procLookupPrivilegeValueW.Addr(), 3, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(luid)))
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func lookupPrivilegeName(systemName string, luid *uint64, buffer *uint16, size *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return
	}
	return _lookupPrivilegeName(_p0, luid, buffer, size)
}

func _lookupPrivilegeName(systemName *uint16, luid *uint64, buffer *uint16, size *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procLookupPrivilegeNameW.Addr(), 4, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(luid)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(size)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func lookupPrivilegeDisplayName(systemName string, name *uint16, buffer *uint16, size *uint32, languageId *uint32) (err error) {
	var _p0 *uint16
	_p0, err = syscall.UTF16PtrFromString(systemName)
	if err != nil {
		return
	}
	return _lookupPrivilegeDisplayName(_p0, name, buffer, size, languageId)
}

func _lookupPrivilegeDisplayName(systemName *uint16, name *uint16, buffer *uint16, size *uint32, languageId *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procLookupPrivilegeDisplayNameW.Addr(), 5, uintptr(unsafe.Pointer(systemName)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(buffer)), uintptr(unsafe.Pointer(size)), uintptr(unsafe.Pointer(languageId)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func backupRead(h syscall.Handle, b []byte, bytesRead *uint32, abort bool, processSecurity bool, context *uintptr) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	var _p1 uint32
	if abort {
		_p1 = 1
	} else {
		_p1 = 0
	}
	var _p2 uint32
	if processSecurity {
		_p2 = 1
	} else {
		_p2 = 0
	}
	r1, _, e1 := syscall.Syscall9(procBackupRead.Addr(), 7, uintptr(h), uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(unsafe.Pointer(bytesRead)), uintptr(_p1), uintptr(_p2), uintptr(unsafe.Pointer(context)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func backupWrite(h syscall.Handle, b []byte, bytesWritten *uint32, abort bool, processSecurity bool, context *uintptr) (err error) {
	var _p0 *byte
	if len(b) > 0 {
		_p0 = &b[0]
	}
	var _p1 uint32
	if abort {
		_p1 = 1
	} else {
		_p1 = 0
	}
	var _p2 uint32
	if processSecurity {
		_p2 = 1
	} else {
		_p2 = 0
	}
	r1, _, e1 := syscall.Syscall9(procBackupWrite.Addr(), 7, uintptr(h), uintptr(unsafe.Pointer(_p0)), uintptr(len(b)), uintptr(unsafe.Pointer(bytesWritten)), uintptr(_p1), uintptr(_p2), uintptr(unsafe.Pointer(context)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
