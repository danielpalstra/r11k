{{- range $key, $value := .Modules }}
mod "{{ $key }}",
  :git => "{{ $value.GitRemote }}"{{ if $value.Ref }},
  :ref => "{{ $value.Ref }}"{{- end }}
{{ end }}
