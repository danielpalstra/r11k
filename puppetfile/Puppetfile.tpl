{{- range $key, $value := .Modules }}
mod "{{ $key }}",
  :git => "{{ $value.GitRemote }}"{{ if $value.Ref }},
  :ref => "{{ $value.Ref }}"{{- end }}{{ if $value.Tag }},
  :tag => "{{ $value.Tag }}"{{- end }}
{{ end }}
