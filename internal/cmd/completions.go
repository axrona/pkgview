package cmd

// Bash completion script for pkgview: caches AUR package names and completes pkgview argument.
const bash = `
__bash_aur_pkg_cache=()

__bash_complete_aur_packages() {
  # If cache is empty, populate it with package names from yay.
  if [ ${#__bash_aur_pkg_cache[@]} -eq 0 ]; then
    mapfile -t __bash_aur_pkg_cache < <(yay -Sl aur | awk '{print $2}')
  fi
  # Print cached package names as completion options.
  printf '%s\n' "${__bash_aur_pkg_cache[@]}"
}

# Register the completion function for pkgview command.
complete -F __bash_complete_aur_packages pkgview
`

// Zsh completion script for pkgview: caches AUR package names and completes pkgview argument.
const zsh = `
typeset -ga __zsh_aur_pkg_cache

__zsh_complete_aur_packages() {
  # If cache array is empty, fill it with package names from yay.
  if (( ${#__zsh_aur_pkg_cache[@]} == 0 )); then
    __zsh_aur_pkg_cache=("${(@f)$(yay -Sl aur | awk '{print $2}')}")
  fi
  # Output cached package names for completion.
  printf '%s\n' "${__zsh_aur_pkg_cache[@]}"
}

# Use compctl to assign completion function to pkgview.
compctl -K __zsh_complete_aur_packages pkgview
`

// Fish completion script for pkgview: caches AUR package names and completes pkgview argument.
const fish = `
function __fish_complete_aur_packages
  # Populate cache if not already set.
  if not set -q __fish_aur_pkg_cache
    set -g __fish_aur_pkg_cache (yay -Sl aur | awk '{print $2}')
  end
  # Print cached package names for completion.
  printf "%s\n" $__fish_aur_pkg_cache
end

# Setup fish completions for pkgview command.
complete -c pkgview -e
complete -c pkgview -f
complete -c pkgview -a "(__fish_complete_aur_packages)"
`

// Map of shell names to their corresponding completion scripts.
var completions = map[string]string{
	"bash": bash,
	"zsh":  zsh,
	"fish": fish,
}
