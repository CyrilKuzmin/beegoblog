<div class="container-fluid text-center">    
  <div class="row content">
    {{ range .Posts }}
    <div class="col-sm-8 text-left"> 
      <h1><a href="/edit?id={{.ID}}">{{ .Title }}</a></h1>
      <p class="post_time">Создан: {{ .CreateTime }}. Последнее изменение: {{ .ModifyTime }}</p>
      {{ .ContentHTML | str2html }}
      <hr>
    </div>
    {{ end }}
  </div>
</div>