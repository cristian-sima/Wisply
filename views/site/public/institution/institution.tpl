<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li class="active">{{ .institution.Name }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <section>
              {{ .institution.URL }}
          </div>
        </div>
      </div>
    </div>
  </div>
  <style>
  .institution-logo {
    font-size: 5em;

  }
  </style>
  <script src="/static/js/admin/institution/list.js"></script>
