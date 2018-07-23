# zsh parameter completion for psm
_psm_zsh_complete()
{
  reply=( $(psm -c "$word") )
}
compctl -K _psm_zsh_complete psm