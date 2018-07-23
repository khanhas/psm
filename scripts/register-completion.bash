# bash parameter completion for psm
_psm_bash_complete()
{
    local word=${COMP_WORDS[COMP_CWORD]}
    COMPREPLY=( $(psm -c "$word") )
}
complete -f -F _psm_bash_complete psm