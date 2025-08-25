# pkgview

🔍 View AUR PKGBUILD files quickly in your terminal.

---

## Why?

The AUR is a user-maintained repository, which sometimes can contain unsafe or malicious software (see [this](https://linuxiac.com/malware-discovered-in-arch-linux-aur-packages) and [this](https://www.bleepingcomputer.com/news/security/malware-found-in-arch-linux-aur-package-repository)). Even if the software is safe, users have the right to know exactly what they're installing.
This simple tool helps you quickly view PKGBUILD files before installing packages.

---

## Installation

### From AUR (recommended)

```bash
yay -S pkgview
```

### Manual installation

```bash
git clone https://github.com/axrona/pkgview.git
cd pkgview
make install
```

--- 

### Usage 

View the PKGBUILD of a package:

```bash
pkgview <package_name>
```

Use a specific editor:

```bash
EDITOR=nano pkgview <package_name>
```

---

### License

This project is licensed under the **GNU GPLv3 License**.

---

<p align="center"><i>"Talk is cheap. Show me the code."</i></p>
