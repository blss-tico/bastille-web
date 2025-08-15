package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func runBastilleCommands(args ...string) (string, error) {
	cmd := exec.Command("bastille", args...)
	log.Println("runBastilleCommands:", cmd)

	result, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("bastille: %s ,failed: %v\n %s", cmd, err, result)
	}

	return string(result), nil
}

func runBastilleCommandAsync(args ...string) (<-chan string, <-chan error) {
	resultChan := make(chan string, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(resultChan)
		defer close(errChan)

		cmd := exec.Command("bastille", args...)
		log.Println("runBastilleCommandAsync:", cmd)

		result, err := cmd.CombinedOutput()
		if err != nil {
			errChan <- fmt.Errorf("bastille: %s, failed: %v\n%s", cmd, err, string(result))
			return
		}

		resultChan <- string(result)
	}()

	return resultChan, errChan
}

type Bastille struct {
}

func (b *Bastille) bastilleVersion() (string, error) {
	args := []string{"-v"}
	return runBastilleCommands(args...)
}

func (b *Bastille) bootstrap(options, releasetemplate, updatearch string) (string, error) {
	args := []string{"bootstrap"}
	if options != "" {
		args = append(args, options)
	}

	if releasetemplate != "" {
		args = append(args, releasetemplate)
	}

	if updatearch != "" {
		args = append(args, updatearch)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) clone(options, target, newname, ip string) (string, error) {
	args := []string{"clone"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if newname != "" {
		args = append(args, newname)
	}

	if ip != "" {
		args = append(args, ip)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) cmd(options, target string, command []string) (string, error) {
	args := []string{"cmd"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	args = append(args, command...)

	res, err := runBastilleCommands(args...)
	return res, err
}

func (b *Bastille) config(options, target, action, property, value string) (string, error) {
	args := []string{"config"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if property != "" {
		args = append(args, property)
	}

	if value != "" {
		args = append(args, value)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) console(options, target, user string) (string, error) {
	args := []string{"console"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if user != "" {
		args = append(args, user)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) convert(options, target, release string) (string, error) {
	args := []string{"convert"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if release != "" {
		args = append(args, release)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) cp(options, target, hostpath, jailpath string) (string, error) {
	args := []string{"convert"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if hostpath != "" {
		args = append(args, hostpath)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) create(options, name, release, ip, iface, ipip, value, vlanid, zfsoptions string) (string, error) {
	args := []string{"create"}

	if options != "" {
		args = append(args, options)
	}

	if name != "" {
		args = append(args, name)
	}

	if release != "" {
		args = append(args, release)
	}

	if ip != "" {
		args = append(args, ip)
	}

	if iface != "" {
		args = append(args, iface)
	}

	if ipip != "" {
		args = append(args, ipip)
	}

	if value != "" {
		args = append(args, value)
	}

	if vlanid != "" {
		args = append(args, vlanid)
	}

	if zfsoptions != "" {
		args = append(args, zfsoptions)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) destroy(options, jailrelease string) (string, error) {
	args := []string{"destroy"}

	if options != "" {
		args = append(args, options)
	}

	if jailrelease != "" {
		args = append(args, jailrelease)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) edit(options, target, file string) (string, error) {
	args := []string{"edit"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if file != "" {
		args = append(args, file)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) etcupdate(options, bootstraptarget, action, release string) (string, error) {
	args := []string{"etcupdate"}

	if options != "" {
		args = append(args, options)
	}

	if bootstraptarget != "" {
		args = append(args, bootstraptarget)
	}

	if action != "" {
		args = append(args, action)
	}

	if release != "" {
		args = append(args, release)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) export(options, target, path string) (string, error) {
	args := []string{"export"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if path != "" {
		args = append(args, path)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) htop(options, target string) (string, error) {
	args := []string{"htop"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) imporT(options, file, release string) (string, error) {
	args := []string{"htop"}

	if options != "" {
		args = append(args, options)
	}

	if file != "" {
		args = append(args, file)
	}

	if release != "" {
		args = append(args, release)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) jcp(options, sourcejail, jailpath, destjail, jailpath2 string) (string, error) {
	args := []string{"jcp"}

	if options != "" {
		args = append(args, options)
	}

	if sourcejail != "" {
		args = append(args, sourcejail)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	if destjail != "" {
		args = append(args, destjail)
	}

	if jailpath2 != "" {
		args = append(args, jailpath2)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) limits(options, target, action, option, value string) (string, error) {
	args := []string{"limits"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if option != "" {
		args = append(args, option)
	}

	if value != "" {
		args = append(args, value)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) list(options, action string) (string, error) {
	args := []string{"list"}

	if options != "" {
		args = append(args, options)
	}

	if action != "" {
		args = append(args, action)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) migrate(options, target, remote string) (string, error) {
	args := []string{"migrate"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if remote != "" {
		args = append(args, remote)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) mount(options, target, hostpath, jailpath, filesystemtype, option, dump, passnumber string) (string, error) {
	args := []string{"mount"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if hostpath != "" {
		args = append(args, hostpath)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	if filesystemtype != "" {
		args = append(args, filesystemtype)
	}

	if option != "" {
		args = append(args, option)
	}

	if dump != "" {
		args = append(args, dump)
	}

	if passnumber != "" {
		args = append(args, passnumber)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) network(options, target, action, iface, ip, vlanid string) (string, error) {
	args := []string{"network"}

	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if iface != "" {
		args = append(args, iface)
	}

	if ip != "" {
		args = append(args, ip)
	}

	if vlanid != "" {
		args = append(args, vlanid)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) pkg(options, target string, arg []string) (string, error) {
	args := []string{"pkg"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	args = append(args, arg...)

	return runBastilleCommands(args...)
}

func (b *Bastille) rcp(options, target, jailpath, hostpath string) (string, error) {
	args := []string{"rcp"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	if hostpath != "" {
		args = append(args, hostpath)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) rdr(options, optionsarg, target, action, hostport, jailport, log, logopts string) (string, error) {
	args := []string{"rdr"}
	if options != "" {
		args = append(args, options)
	}

	if optionsarg != "" {
		args = append(args, optionsarg)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if hostport != "" {
		args = append(args, hostport)
	}

	if jailport != "" {
		args = append(args, jailport)
	}

	if log != "" {
		args = append(args, log)
	}

	if logopts != "" {
		args = append(args, logopts)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) rename(options, target, newname string) (string, error) {
	args := []string{"rename"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if newname != "" {
		args = append(args, newname)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) restart(options, target, value string) (string, error) {
	args := []string{"restart"}
	if options != "" {
		args = append(args, options)
	}

	if value != "" {
		args = append(args, value)
	}

	if target != "" {
		args = append(args, target)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) service(options, target, servicename, arg string) (string, error) {
	args := []string{"service"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if servicename != "" {
		args = append(args, servicename)
	}

	if arg != "" {
		var formatedArgs []string
		formatedArgs = append(formatedArgs, "'")
		formatedArgs = append(formatedArgs, arg)
		formatedArgs = append(formatedArgs, "'")
		result := strings.Join(formatedArgs, " ")
		args = append(args, result)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) setup(options, action string) (string, error) {
	args := []string{"setup"}
	if options != "" {
		args = append(args, options)
	}

	if action != "" {
		args = append(args, action)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) start(options, target, value string) (string, error) {
	args := []string{"start"}
	if options != "" {
		args = append(args, options)
	}

	if value != "" {
		args = append(args, value)
	}

	if target != "" {
		args = append(args, target)
	}

	//return runBastilleCommands(args...)
	resultChan, errChan := runBastilleCommandAsync(args...)
	select {
	case res := <-resultChan:
		return res, nil
	case err := <-errChan:
		return "", err
	}
}

func (b *Bastille) stop(options, target string) (string, error) {
	args := []string{"stop"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	// return runBastilleCommands(args...)
	resultChan, errChan := runBastilleCommandAsync(args...)
	select {
	case res := <-resultChan:
		return res, nil
	case err := <-errChan:
		return "", err
	}
}

func (b *Bastille) sysrc(options, target, arg string) (string, error) {
	args := []string{"sysrc"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if arg != "" {
		args = append(args, arg)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) tags(options, target, action, tags string) (string, error) {
	args := []string{"tags"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if tags != "" {
		args = append(args, tags)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) template(options, target, action, template string) (string, error) {
	args := []string{"template"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if template != "" {
		args = append(args, template)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) top(options, target string) (string, error) {
	args := []string{"top"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) umount(options, target, jailpath string) (string, error) {
	args := []string{"umount"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) update(options, target string) (string, error) {
	args := []string{"update"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) upgrade(options, target, action string) (string, error) {
	args := []string{"upgrade"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) verify(options, action string) (string, error) {
	args := []string{"verify"}
	if options != "" {
		args = append(args, options)
	}

	if action != "" {
		args = append(args, action)
	}

	return runBastilleCommands(args...)
}

func (b *Bastille) zfs(options, target, action, tag, key, value, pooldataset, jailpath string) (string, error) {
	args := []string{"zfs"}
	if options != "" {
		args = append(args, options)
	}

	if target != "" {
		args = append(args, target)
	}

	if action != "" {
		args = append(args, action)
	}

	if tag != "" {
		args = append(args, tag)
	}

	if key != "" {
		args = append(args, key)
	}

	if value != "" {
		args = append(args, value)
	}

	if pooldataset != "" {
		args = append(args, pooldataset)
	}

	if jailpath != "" {
		args = append(args, jailpath)
	}

	return runBastilleCommands(args...)
}
