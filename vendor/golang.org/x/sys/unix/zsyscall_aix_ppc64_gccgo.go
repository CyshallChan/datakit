// go run mksyscall_aix_ppc64.go -aix -tags aix,ppc64 syscall_aix.go syscall_aix_ppc64.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build aix,ppc64
// +build gccgo

package unix

/*
#include <stdint.h>
int utimes(uintptr_t, uintptr_t);
int utimensat(int, uintptr_t, uintptr_t, int);
int getcwd(uintptr_t, size_t);
int accept(int, uintptr_t, uintptr_t);
int getdirent(int, uintptr_t, size_t);
int wait4(int, uintptr_t, int, uintptr_t);
int ioctl(int, int, uintptr_t);
int fcntl(uintptr_t, int, uintptr_t);
int acct(uintptr_t);
int chdir(uintptr_t);
int chroot(uintptr_t);
int close(int);
int dup(int);
void exit(int);
int faccessat(int, uintptr_t, unsigned int, int);
int fchdir(int);
int fchmod(int, unsigned int);
int fchmodat(int, uintptr_t, unsigned int, int);
int fchownat(int, uintptr_t, int, int, int);
int fdatasync(int);
int fsync(int);
int getpgid(int);
int getpgrp();
int getpid();
int getppid();
int getpriority(int, int);
int getrusage(int, uintptr_t);
int getsid(int);
int kill(int, int);
int syslog(int, uintptr_t, size_t);
int mkdir(int, uintptr_t, unsigned int);
int mkdirat(int, uintptr_t, unsigned int);
int mkfifo(uintptr_t, unsigned int);
int mknod(uintptr_t, unsigned int, int);
int mknodat(int, uintptr_t, unsigned int, int);
int nanosleep(uintptr_t, uintptr_t);
int open64(uintptr_t, int, unsigned int);
int openat(int, uintptr_t, int, unsigned int);
int read(int, uintptr_t, size_t);
int readlink(uintptr_t, uintptr_t, size_t);
int renameat(int, uintptr_t, int, uintptr_t);
int setdomainname(uintptr_t, size_t);
int sethostname(uintptr_t, size_t);
int setpgid(int, int);
int setsid();
int settimeofday(uintptr_t);
int setuid(int);
int setgid(int);
int setpriority(int, int, int);
int statx(int, uintptr_t, int, int, uintptr_t);
int sync();
uintptr_t times(uintptr_t);
int umask(int);
int uname(uintptr_t);
int unlink(uintptr_t);
int unlinkat(int, uintptr_t, int);
int ustat(int, uintptr_t);
int write(int, uintptr_t, size_t);
int dup2(int, int);
int posix_fadvise64(int, long long, long long, int);
int fchown(int, int, int);
int fstat(int, uintptr_t);
int fstatat(int, uintptr_t, uintptr_t, int);
int fstatfs(int, uintptr_t);
int ftruncate(int, long long);
int getegid();
int geteuid();
int getgid();
int getuid();
int lchown(uintptr_t, int, int);
int listen(int, int);
int lstat(uintptr_t, uintptr_t);
int pause();
int pread64(int, uintptr_t, size_t, long long);
int pwrite64(int, uintptr_t, size_t, long long);
#define c_select select
int select(int, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
int pselect(int, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);
int setregid(int, int);
int setreuid(int, int);
int shutdown(int, int);
long long splice(int, uintptr_t, int, uintptr_t, int, int);
int stat(uintptr_t, uintptr_t);
int statfs(uintptr_t, uintptr_t);
int truncate(uintptr_t, long long);
int bind(int, uintptr_t, uintptr_t);
int connect(int, uintptr_t, uintptr_t);
int getgroups(int, uintptr_t);
int setgroups(int, uintptr_t);
int getsockopt(int, int, int, uintptr_t, uintptr_t);
int setsockopt(int, int, int, uintptr_t, uintptr_t);
int socket(int, int, int);
int socketpair(int, int, int, uintptr_t);
int getpeername(int, uintptr_t, uintptr_t);
int getsockname(int, uintptr_t, uintptr_t);
int recvfrom(int, uintptr_t, size_t, int, uintptr_t, uintptr_t);
int sendto(int, uintptr_t, size_t, int, uintptr_t, uintptr_t);
int nrecvmsg(int, uintptr_t, int);
int nsendmsg(int, uintptr_t, int);
int munmap(uintptr_t, uintptr_t);
int madvise(uintptr_t, size_t, int);
int mprotect(uintptr_t, size_t, int);
int mlock(uintptr_t, size_t);
int mlockall(int);
int msync(uintptr_t, size_t, int);
int munlock(uintptr_t, size_t);
int munlockall();
int pipe(uintptr_t);
int poll(uintptr_t, int, int);
int gettimeofday(uintptr_t, uintptr_t);
int time(uintptr_t);
int utime(uintptr_t, uintptr_t);
unsigned long long getsystemcfg(int);
int umount(uintptr_t);
int getrlimit(int, uintptr_t);
int setrlimit(int, uintptr_t);
long long lseek(int, long long, int);
uintptr_t mmap64(uintptr_t, uintptr_t, int, int, int, long long);

*/
import "C"
import (
	"syscall"
)

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callutimes(_p0 uintptr, times uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.utimes(C.uintptr_t(_p0), C.uintptr_t(times)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callutimensat(dirfd int, _p0 uintptr, times uintptr, flag int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.utimensat(C.int(dirfd), C.uintptr_t(_p0), C.uintptr_t(times), C.int(flag)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetcwd(_p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getcwd(C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callaccept(s int, rsa uintptr, addrlen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.accept(C.int(s), C.uintptr_t(rsa), C.uintptr_t(addrlen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetdirent(fd int, _p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getdirent(C.int(fd), C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callwait4(pid int, status uintptr, options int, rusage uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.wait4(C.int(pid), C.uintptr_t(status), C.int(options), C.uintptr_t(rusage)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callioctl(fd int, req int, arg uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.ioctl(C.int(fd), C.int(req), C.uintptr_t(arg)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfcntl(fd uintptr, cmd int, arg uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fcntl(C.uintptr_t(fd), C.int(cmd), C.uintptr_t(arg)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callacct(_p0 uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.acct(C.uintptr_t(_p0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callchdir(_p0 uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.chdir(C.uintptr_t(_p0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callchroot(_p0 uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.chroot(C.uintptr_t(_p0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callclose(fd int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.close(C.int(fd)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calldup(oldfd int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.dup(C.int(oldfd)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callexit(code int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.exit(C.int(code)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfaccessat(dirfd int, _p0 uintptr, mode uint32, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.faccessat(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfchdir(fd int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fchdir(C.int(fd)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfchmod(fd int, mode uint32) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fchmod(C.int(fd), C.uint(mode)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfchmodat(dirfd int, _p0 uintptr, mode uint32, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fchmodat(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfchownat(dirfd int, _p0 uintptr, uid int, gid int, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fchownat(C.int(dirfd), C.uintptr_t(_p0), C.int(uid), C.int(gid), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfdatasync(fd int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fdatasync(C.int(fd)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfsync(fd int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fsync(C.int(fd)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetpgid(pid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getpgid(C.int(pid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetpgrp() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getpgrp())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetpid() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getpid())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetppid() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getppid())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetpriority(which int, who int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getpriority(C.int(which), C.int(who)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetrusage(who int, rusage uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getrusage(C.int(who), C.uintptr_t(rusage)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetsid(pid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getsid(C.int(pid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callkill(pid int, sig int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.kill(C.int(pid), C.int(sig)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsyslog(typ int, _p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.syslog(C.int(typ), C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmkdir(dirfd int, _p0 uintptr, mode uint32) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mkdir(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmkdirat(dirfd int, _p0 uintptr, mode uint32) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mkdirat(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmkfifo(_p0 uintptr, mode uint32) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mkfifo(C.uintptr_t(_p0), C.uint(mode)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmknod(_p0 uintptr, mode uint32, dev int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mknod(C.uintptr_t(_p0), C.uint(mode), C.int(dev)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmknodat(dirfd int, _p0 uintptr, mode uint32, dev int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mknodat(C.int(dirfd), C.uintptr_t(_p0), C.uint(mode), C.int(dev)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callnanosleep(time uintptr, leftover uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.nanosleep(C.uintptr_t(time), C.uintptr_t(leftover)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callopen64(_p0 uintptr, mode int, perm uint32) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.open64(C.uintptr_t(_p0), C.int(mode), C.uint(perm)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callopenat(dirfd int, _p0 uintptr, flags int, mode uint32) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.openat(C.int(dirfd), C.uintptr_t(_p0), C.int(flags), C.uint(mode)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callread(fd int, _p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.read(C.int(fd), C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callreadlink(_p0 uintptr, _p1 uintptr, _lenp1 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.readlink(C.uintptr_t(_p0), C.uintptr_t(_p1), C.size_t(_lenp1)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callrenameat(olddirfd int, _p0 uintptr, newdirfd int, _p1 uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.renameat(C.int(olddirfd), C.uintptr_t(_p0), C.int(newdirfd), C.uintptr_t(_p1)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetdomainname(_p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setdomainname(C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsethostname(_p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.sethostname(C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetpgid(pid int, pgid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setpgid(C.int(pid), C.int(pgid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetsid() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setsid())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsettimeofday(tv uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.settimeofday(C.uintptr_t(tv)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetuid(uid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setuid(C.int(uid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetgid(uid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setgid(C.int(uid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetpriority(which int, who int, prio int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setpriority(C.int(which), C.int(who), C.int(prio)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callstatx(dirfd int, _p0 uintptr, flags int, mask int, stat uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.statx(C.int(dirfd), C.uintptr_t(_p0), C.int(flags), C.int(mask), C.uintptr_t(stat)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsync() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.sync())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calltimes(tms uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.times(C.uintptr_t(tms)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callumask(mask int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.umask(C.int(mask)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calluname(buf uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.uname(C.uintptr_t(buf)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callunlink(_p0 uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.unlink(C.uintptr_t(_p0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callunlinkat(dirfd int, _p0 uintptr, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.unlinkat(C.int(dirfd), C.uintptr_t(_p0), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callustat(dev int, ubuf uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.ustat(C.int(dev), C.uintptr_t(ubuf)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callwrite(fd int, _p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.write(C.int(fd), C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calldup2(oldfd int, newfd int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.dup2(C.int(oldfd), C.int(newfd)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callposix_fadvise64(fd int, offset int64, length int64, advice int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.posix_fadvise64(C.int(fd), C.longlong(offset), C.longlong(length), C.int(advice)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfchown(fd int, uid int, gid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fchown(C.int(fd), C.int(uid), C.int(gid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfstat(fd int, stat uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fstat(C.int(fd), C.uintptr_t(stat)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfstatat(dirfd int, _p0 uintptr, stat uintptr, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fstatat(C.int(dirfd), C.uintptr_t(_p0), C.uintptr_t(stat), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callfstatfs(fd int, buf uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.fstatfs(C.int(fd), C.uintptr_t(buf)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callftruncate(fd int, length int64) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.ftruncate(C.int(fd), C.longlong(length)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetegid() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getegid())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgeteuid() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.geteuid())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetgid() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getgid())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetuid() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getuid())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calllchown(_p0 uintptr, uid int, gid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.lchown(C.uintptr_t(_p0), C.int(uid), C.int(gid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calllisten(s int, n int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.listen(C.int(s), C.int(n)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calllstat(_p0 uintptr, stat uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.lstat(C.uintptr_t(_p0), C.uintptr_t(stat)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callpause() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.pause())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callpread64(fd int, _p0 uintptr, _lenp0 int, offset int64) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.pread64(C.int(fd), C.uintptr_t(_p0), C.size_t(_lenp0), C.longlong(offset)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callpwrite64(fd int, _p0 uintptr, _lenp0 int, offset int64) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.pwrite64(C.int(fd), C.uintptr_t(_p0), C.size_t(_lenp0), C.longlong(offset)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callselect(nfd int, r uintptr, w uintptr, e uintptr, timeout uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.c_select(C.int(nfd), C.uintptr_t(r), C.uintptr_t(w), C.uintptr_t(e), C.uintptr_t(timeout)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callpselect(nfd int, r uintptr, w uintptr, e uintptr, timeout uintptr, sigmask uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.pselect(C.int(nfd), C.uintptr_t(r), C.uintptr_t(w), C.uintptr_t(e), C.uintptr_t(timeout), C.uintptr_t(sigmask)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetregid(rgid int, egid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setregid(C.int(rgid), C.int(egid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetreuid(ruid int, euid int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setreuid(C.int(ruid), C.int(euid)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callshutdown(fd int, how int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.shutdown(C.int(fd), C.int(how)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsplice(rfd int, roff uintptr, wfd int, woff uintptr, len int, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.splice(C.int(rfd), C.uintptr_t(roff), C.int(wfd), C.uintptr_t(woff), C.int(len), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callstat(_p0 uintptr, statptr uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.stat(C.uintptr_t(_p0), C.uintptr_t(statptr)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callstatfs(_p0 uintptr, buf uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.statfs(C.uintptr_t(_p0), C.uintptr_t(buf)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calltruncate(_p0 uintptr, length int64) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.truncate(C.uintptr_t(_p0), C.longlong(length)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callbind(s int, addr uintptr, addrlen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.bind(C.int(s), C.uintptr_t(addr), C.uintptr_t(addrlen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callconnect(s int, addr uintptr, addrlen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.connect(C.int(s), C.uintptr_t(addr), C.uintptr_t(addrlen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetgroups(n int, list uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getgroups(C.int(n), C.uintptr_t(list)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetgroups(n int, list uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setgroups(C.int(n), C.uintptr_t(list)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetsockopt(s int, level int, name int, val uintptr, vallen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getsockopt(C.int(s), C.int(level), C.int(name), C.uintptr_t(val), C.uintptr_t(vallen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetsockopt(s int, level int, name int, val uintptr, vallen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setsockopt(C.int(s), C.int(level), C.int(name), C.uintptr_t(val), C.uintptr_t(vallen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsocket(domain int, typ int, proto int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.socket(C.int(domain), C.int(typ), C.int(proto)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsocketpair(domain int, typ int, proto int, fd uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.socketpair(C.int(domain), C.int(typ), C.int(proto), C.uintptr_t(fd)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetpeername(fd int, rsa uintptr, addrlen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getpeername(C.int(fd), C.uintptr_t(rsa), C.uintptr_t(addrlen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetsockname(fd int, rsa uintptr, addrlen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getsockname(C.int(fd), C.uintptr_t(rsa), C.uintptr_t(addrlen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callrecvfrom(fd int, _p0 uintptr, _lenp0 int, flags int, from uintptr, fromlen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.recvfrom(C.int(fd), C.uintptr_t(_p0), C.size_t(_lenp0), C.int(flags), C.uintptr_t(from), C.uintptr_t(fromlen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsendto(s int, _p0 uintptr, _lenp0 int, flags int, to uintptr, addrlen uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.sendto(C.int(s), C.uintptr_t(_p0), C.size_t(_lenp0), C.int(flags), C.uintptr_t(to), C.uintptr_t(addrlen)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callnrecvmsg(s int, msg uintptr, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.nrecvmsg(C.int(s), C.uintptr_t(msg), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callnsendmsg(s int, msg uintptr, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.nsendmsg(C.int(s), C.uintptr_t(msg), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmunmap(addr uintptr, length uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.munmap(C.uintptr_t(addr), C.uintptr_t(length)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmadvise(_p0 uintptr, _lenp0 int, advice int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.madvise(C.uintptr_t(_p0), C.size_t(_lenp0), C.int(advice)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmprotect(_p0 uintptr, _lenp0 int, prot int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mprotect(C.uintptr_t(_p0), C.size_t(_lenp0), C.int(prot)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmlock(_p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mlock(C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmlockall(flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mlockall(C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmsync(_p0 uintptr, _lenp0 int, flags int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.msync(C.uintptr_t(_p0), C.size_t(_lenp0), C.int(flags)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmunlock(_p0 uintptr, _lenp0 int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.munlock(C.uintptr_t(_p0), C.size_t(_lenp0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmunlockall() (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.munlockall())
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callpipe(p uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.pipe(C.uintptr_t(p)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callpoll(fds uintptr, nfds int, timeout int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.poll(C.uintptr_t(fds), C.int(nfds), C.int(timeout)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgettimeofday(tv uintptr, tzp uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.gettimeofday(C.uintptr_t(tv), C.uintptr_t(tzp)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calltime(t uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.time(C.uintptr_t(t)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callutime(_p0 uintptr, buf uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.utime(C.uintptr_t(_p0), C.uintptr_t(buf)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetsystemcfg(label int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getsystemcfg(C.int(label)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callumount(_p0 uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.umount(C.uintptr_t(_p0)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callgetrlimit(resource int, rlim uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.getrlimit(C.int(resource), C.uintptr_t(rlim)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callsetrlimit(resource int, rlim uintptr) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.setrlimit(C.int(resource), C.uintptr_t(rlim)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func calllseek(fd int, offset int64, whence int) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.lseek(C.int(fd), C.longlong(offset), C.int(whence)))
	e1 = syscall.GetErrno()
	return
}

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func callmmap64(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (r1 uintptr, e1 Errno) {
	r1 = uintptr(C.mmap64(C.uintptr_t(addr), C.uintptr_t(length), C.int(prot), C.int(flags), C.int(fd), C.longlong(offset)))
	e1 = syscall.GetErrno()
	return
}
