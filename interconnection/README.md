# Interconnection 🔗
Is a set of different services that help with the exchanges of sensitive information between applications.

## Exchange Methods
There are many ways to exchange internally in a system but not all are secure. It is an option to print out the
secrets that `doorman` receives from one of the *clerks* (the name of the plugins for `doorman`). It can be pass
securely by piping it to another terminal application.

But there might be situations where this option is the least secure one. Also, *clerks* need to be able to exchange with
`doorman` in a secure way. So for those reasons three methods are available:

- 🪟 [Named Pipes](https://learn.microsoft.com/en-us/windows/win32/ipc/named-pipes) - Only for Windows
- 🔌 [Unix Domain Sockets](https://man7.org/linux/man-pages/man7/unix.7.html) - Available for Mac OS X 🍎 and Linux 🐧
- 🚌 [D-Bus](https://www.freedesktop.org/wiki/Software/dbus/) - Only for Linux 🐧

### Security comparison

| Threat Model	                         | 🚌 D-Bus	                                | 🔌 Unix Domain Sockets                 |
|---------------------------------------|------------------------------------------|----------------------------------------|
| Prevent unauthorized access	          | ✅ Built-in access policies	              | ✅ File permissions (chmod)             |
| Prevent privilege escalation	         | ✅ Restricted to specific users/apps	     | ✅ Restricted by filesystem permissions |
| Prevent data interception	            | ✅ Encrypted if configured	               | 🚫 No encryption (local IPC only)      |
| Prevent unauthorized service control	 | ✅ Supports fine-grained authentication	  | 🚫 No built-in authentication          |
