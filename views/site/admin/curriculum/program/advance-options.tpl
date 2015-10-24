<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/curriculum">Curriculum</a></li>
      <li><a href="/admin/curriculum/programs/{{ .program.GetName }}">{{ .program.GetName }}</a></li>
      <li class="active">Advance options</li>
    </ul></div>
    <div class="panel-body">
      <div class="row" >
        <div class="col-md-12">
          <div>
            <h2>Modify</h2>
            <div class="well">
              You can use this option to modify the details of the program. <br/>
              <a href="/admin/curriculum/programs/{{ .program.GetID }}/advance-options/modify" class="btn btn-primary">Modify program</a>
            </div>
          </div>
          <div>
            <h2>Delete</h2>
            <div class="well">
              <h5><span class="text-warning"><span class="glyphicon glyphicon-warning-sign"></span> Please note that all the analyse for this program can not be recovered.</span></h5><br />
              <a data-id="{{ .program.GetID }}" data-name="" href="#" class="btn btn-danger deleteProgramButton">Delete program and entire data related to it from Wisply</a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script src="/static/js/admin/repository/list.js"></script>
