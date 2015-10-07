<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/log">Event log</a></li>
      <li><a href="/admin/log/process/{{ .process.Action.ID }}">Process #{{ .process.Action.ID }}</a></li>
      <li class="active">Advance options</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="row" >
      <div class="col-md-12">
        <div>
          <h2>Delete</h2>
          <div class="well">
            Please note that the process can not be recovered.
            <a data-id="" data-name="" href="#" class="btn btn-danger deleteProcessButton">Delete process from Wisply</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
