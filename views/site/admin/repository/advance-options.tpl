<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      <li><a href="/admin/repositories/repository/{{ .repository.ID }}">{{ .repository.Name }}</a></li>
      <li class="active">Advance options</li>
    </ul></div>
    <div class="panel-body">
        <div class="row" >
          <div class="col-md-12">
            <div>
              <h2>Modify</h2>
              <div class="well">
                You can use this option to modify the details of the repository. It is impossible to modify the Base URL
                <a href="/admin/repositories/modify/{{ .repository.ID }}" class="btn btn-primary">Modify repository</a>
              </div>
            </div>
            <div>
              <h2>Delete</h2>
              <div class="well">
                Please note that the repository can not be recovered.
                <a class="deleteRepositoryButton" data-id="" data-name="" href="#" class="btn btn-danger">Delete repository</a>
              </div>
            </div>
          </div>
        </div>
    </div>
  </div>
  <style>
  .big-number {
    font-size: 30px;
  }
  </style>
  <script src="/static/js/admin/repository/list.js"></script>
  <script>
  $(document).ready(function(){

  });
  </script>
