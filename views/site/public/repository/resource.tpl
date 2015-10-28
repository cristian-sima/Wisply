<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li><a href="/repository/{{ .repository.ID }}">{{ .repository.Name }}</a></li>
            <li class="active">{{ .record.Keys.GetTitle }}</li>
          </ul>
        </div>
        <div class="panel-body">
            <div class="top-info">
            </div>

            <div class="content-info">
              <div class="embed-responsive embed-responsive-16by9">
                  <base target="_blank" />
              <iframe id="the-iframe" sandbox class="embed-responsive-item the-iframe" src="http://facebook.com">

              </iframe>
            </div>
                {{ .record }}
            </div>
        </div>
      </div>
    </div>
  </div>
</div>
<div>
  <script>
  $.ajax({
  url: 'http://facebook.com',
  type: 'GET',
  success: function(res) {
    $("#the-iframe").html(res.responseText)
  }
});
  </script>
<style scoped>
.the-iframe {

}
</style>
</div>
