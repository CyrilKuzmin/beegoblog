<script type="text/javascript" src="/static/js/confirm.js"></script>
<script type="text/javascript" src="https://cdn.tinymce.com/4/tinymce.min.js"></script>
<script type="application/x-javascript">
tinymce.init({
  selector:'#TypeHere',
  plugins: "image code"
});
</script>
<div class="container-fluid text-center">    
  <div class="row content">
    <div class="col-sm-8 text-left"> 
      <h1>Добавление нового поста</h1>
      <hr>
      <form role="form" method="POST" action="/savepost">
      <div class="col-xs-8">
            <input type="hidden" name="id" value="{{.Post.ID}}"/>
            <div class="form-group">
                <label>Заголовок</label>
                <input type="text" class="form-control" id="title" name="title" value="{{.Post.Title}}" required/>
            </div>
            <div class="form-group">
                <label>Содержимое</label>
                <textarea class="form-control" id="TypeHere" name="editor" rows="10" data-ng-model="text">{{.Post.ContentHTML}}</textarea>
            </div>
            <div>
              <button type="submit" class="btn btn-success">Отправить</button>
              {{ if .Post.ID }}
              <button type="reset" class="btn btn-default">Отменить изменения</button>
              {{ end }}
        
          {{ if .Post.ID }}
          <a href="/delete?id={{.Post.ID}}" class="btn btn-danger"  onclick="return confirmRemoval();">Удалить</a>
          {{ end }}
          </div>
        </div>
      </form>
      </div>
    </div>
  </div>
</div>