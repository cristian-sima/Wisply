<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      <li><a href="/admin/repositories/{{ .repository.ID }}">{{ .repository.Name }}</a></li>
      <li class="active">Advance options</li>
    </ul></div>
    <div class="panel-body">
      <div class="row" >
        <div class="col-md-12">
          <div>
            <h2>Modify</h2>
            <div class="well">
              You can use this option to modify the details of the repository. <br/>
              <a href="/admin/repositories/{{ .repository.ID }}/advance-options/modify" class="btn btn-primary">Modify repository</a>
            </div>
          </div>
          <div>
            <h2>Filter</h2>
            <div class="well">
              You can filter what input you harvest or process.
               <br/>
              <a href="/admin/repositories/{{ .repository.ID }}/advance-options/modify/filter" class="btn btn-primary">Modify filter</a>
            </div>
          </div>
          <div>
            <h2>Clear metadata</h2>
            <div class="well">
              This option can be used to clear all the records, collections, formats, identifiers, emails and the identification details.
              <h5><span class="text-warning"><span class="glyphicon glyphicon-warning-sign"></span> Please note that the items can not be recovered.</span></h5><br />
              <a data-id="{{ .repository.ID }}" data-name="" href="#" class="btn btn-danger emptyRepositoryButton">Clear metadata from {{ .repository.Name }}</a>
            </div>
          </div>
          <div>
            <h2>Delete</h2>
            <div class="well">
              <h5><span class="text-warning"><span class="glyphicon glyphicon-warning-sign"></span> Please note that the repository can not be recovered.</span></h5><br />
              <a data-id="{{ .repository.ID }}" data-name="" href="#" class="btn btn-danger deleteRepositoryButton">Delete {{ .repository.Name }}</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div>
    <style scoped>
    .big-number {
      font-size: 30px;
    }
    </style>
  </div>
  <script src="/static/js/admin/repositories/repository/list.js"></script>
  <script>
  $(document).ready(function(){
    var module = wisply.getModule("admin-repositories-list"),
      manager = new module.Manager();
      manager.init();
  });
  </script>
