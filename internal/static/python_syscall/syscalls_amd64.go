//go:build linux && amd64

package python_syscall

import "syscall"

const (
	SYS_GETRANDOM = 318
	SYS_RSEQ      = 334
	SYS_SENDMMSG  = 307
)

var ALLOW_SYSCALLS = []int{
	// file io
	syscall.SYS_NEWFSTATAT, syscall.SYS_IOCTL, syscall.SYS_LSEEK, syscall.SYS_GETDENTS64,
	syscall.SYS_WRITE, syscall.SYS_CLOSE, syscall.SYS_OPENAT, syscall.SYS_READ,
	// thread
	syscall.SYS_FUTEX,
	// memory
	syscall.SYS_MMAP, syscall.SYS_BRK, syscall.SYS_MPROTECT, syscall.SYS_MUNMAP, syscall.SYS_RT_SIGRETURN,
	syscall.SYS_MREMAP,

	// user/group
	syscall.SYS_SETUID, syscall.SYS_SETGID, syscall.SYS_GETUID,
	// process
	syscall.SYS_GETPID, syscall.SYS_GETPPID, syscall.SYS_GETTID,
	syscall.SYS_EXIT, syscall.SYS_EXIT_GROUP,
	syscall.SYS_TGKILL, syscall.SYS_RT_SIGACTION, syscall.SYS_IOCTL,
	syscall.SYS_SCHED_YIELD,
	syscall.SYS_SET_ROBUST_LIST, syscall.SYS_GET_ROBUST_LIST, SYS_RSEQ,

	// time
	syscall.SYS_CLOCK_GETTIME, syscall.SYS_GETTIMEOFDAY, syscall.SYS_NANOSLEEP,
	syscall.SYS_EPOLL_CREATE1,
	syscall.SYS_EPOLL_CTL, syscall.SYS_CLOCK_NANOSLEEP, syscall.SYS_PSELECT6,
	syscall.SYS_TIME,

	syscall.SYS_RT_SIGPROCMASK, syscall.SYS_SIGALTSTACK, SYS_GETRANDOM,

	// ── new additions from strace -c ──

	// basic file operations
	syscall.SYS_STAT, syscall.SYS_FSTAT, syscall.SYS_LSTAT,
	syscall.SYS_PREAD64, syscall.SYS_ACCESS, syscall.SYS_READLINK,

	// more IO variants
	syscall.SYS_FCNTL, syscall.SYS_GETCWD, syscall.SYS_MKDIR,

	// memory & mapping
	// (mmap, brk, mprotect, munmap already above)

	// control & device
	// (ioctl already above)

	// signals & threading
	// (futex, rt_sigaction, rt_sigprocmask already above)
	syscall.SYS_SET_TID_ADDRESS,

	// randomness
	// (getrandom / SYS_GETRANDOM already above)

	// architecture-specific
	syscall.SYS_ARCH_PRCTL,

	// resource limits
	syscall.SYS_PRLIMIT64,

	// process management
	syscall.SYS_UNAME, syscall.SYS_CLONE, syscall.SYS_EXECVE, syscall.SYS_SYSINFO,
	syscall.SYS_SCHED_GETAFFINITY, syscall.SYS_MBIND,
}


var ALLOW_ERROR_SYSCALLS = []int{
	syscall.SYS_CLONE,
	syscall.SYS_MKDIRAT,
	syscall.SYS_MKDIR,
}

var ALLOW_NETWORK_SYSCALLS = []int{
	syscall.SYS_SOCKET, syscall.SYS_CONNECT, syscall.SYS_BIND, syscall.SYS_LISTEN, syscall.SYS_ACCEPT, syscall.SYS_SENDTO, syscall.SYS_RECVFROM,
	syscall.SYS_GETSOCKNAME, syscall.SYS_RECVMSG, syscall.SYS_GETPEERNAME, syscall.SYS_SETSOCKOPT, syscall.SYS_PPOLL, syscall.SYS_UNAME,
	syscall.SYS_SENDMSG, SYS_SENDMMSG, syscall.SYS_GETSOCKOPT,
	syscall.SYS_FSTAT, syscall.SYS_FCNTL, syscall.SYS_FSTATFS, syscall.SYS_POLL, syscall.SYS_EPOLL_PWAIT,
}
