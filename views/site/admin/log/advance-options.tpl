<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/log">Event log</a></li>
      <li class="active">Advance options</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="row" >
      <div class="col-md-12">
        <div>
          <h2>Delete the entire log</h2>
          <div class="well">
            <h5><span class="text-warning"><span class="glyphicon glyphicon-warning-sign"></span> Please note that the log can not be recovered.</span></h5><br />
            <a  href="#" class="btn btn-danger deleteLogButton">Delete entire log for ever from Wisply</a>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<script src="/static/js/admin/log/log-advance-options.js"></script>
<script>
$(document).ready(function(){
    var module = wisply.getModule("admin-log-advance-options"),
      manager = new module.Manager();
      manager.init();
});
</script>
