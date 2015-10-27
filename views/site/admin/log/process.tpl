<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/log">Event log</a></li>
      <li class="active">Process #{{ .process.Action.ID }}</li>
    </ul>
  </div>
  <div class="panel-body">
    <div>
      <span class="text-warning warning-notice">
        <span class="glyphicon glyphicon-warning-sign"></span>
        This page is not live updated.
      </span>
    </div>
    <h2>
      Process #{{ .process.Action.ID }}
    </h2>
    <div class="row">
      <div class="col-lg-4 col-md-4 col-sm-4">
        <table class="table">
          <tbody>
            <tr>
              <td>Type</td>
              <td><strong>{{ .process.Action.Content }}</strong></td>
            </tr>
            <tr>
              <td>Repository</td>
              <td><strong><a href="/admin/repositories/repository/{{ .harvestProcess.GetRepository.ID }}">{{ .harvestProcess.GetRepository.Name }}</a></strong></td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="col-lg-4 col-md-4 col-sm-4">
        <table class="table">
          <tbody>
            <tr>
              <td><span class="glyphicon glyphicon-calendar"></span> Start:</td>
              <td>{{ .process.GetStartDate }}</td>
            </tr>
            <tr>
              <td><span class="glyphicon glyphicon-calendar"></span> Finish:</td>
              <td><strong>{{ .process.GetEndDate }}</strong></td>
            </tr>
          </tbody>
        </table>
      </div>
      <div class="col-lg-4 col-md-4 col-sm-4">
        <table class="table">
          <tbody>
            <tr>
              <td>
                {{ if .process.IsSuspended }}
                <span class="label label-warning">Suspended</span>
                {{ else }}
                {{ if .process.Action.IsRunning }}
                <span class="label label-info">Working</span>
                {{ else }}
                <span class="label label-success">Finished</span>
                {{ end }}
                {{ end }}
              </td>
              <td>
                <strong>
                  {{ if .process.Action.IsRunning }}
                  <a href="/admin/log/process/{{.process.Action.ID}}/operation/{{.operation.ID}}">{{ .operation.Content }}</a>
                  {{ else }}
                  <span class="glyphicon glyphicon-time"></span> {{ .process.GetDuration }}
                  {{ end }}
                </strong>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
    {{ if .process.IsSuspended }}
    <div class="well">
      <h3>Warning</h3>
      <span class="text-warning">This process has been forced to stop, because of a problem.</span> <br />
      We need your action: <br /><br />
      <a href="#" id="recover-button" data-id="{{.harvestProcess.HarvestID}}" class="btn btn-primary">Try again</a>
      <a href="#" id="finish-button" data-id="{{.harvestProcess.HarvestID}}" class="btn btn-primary">Just Finish</a>
    </div>
    {{ end }}
    <div class="no-print">
      <a class="btn btn-primary" href="/admin/log/process/{{ .process.Action.ID }}/history#history">Show history</a>
      <a class="btn btn-primary" href="/admin/log/process/{{ .process.Action.ID }}/advance-options">Advance options</a>
      <a class="btn btn-primary printPage" ><span class="glyphicon glyphicon-print"></span> Print</a>
    </div>
    <br />
    {{ $len := .operations | len }}
    {{ if eq $len 0 }}
    <div class="text-center">
      There are no operations for this process
    </div>
    {{ else }}
    <div class="print-div">
      <div class="table-responsive">
        <table id="list-operations" class="table table-bordered table-condensed">
          <thead>
            <tr>
              <th>#Operation</th>
              <th>Content</th>
              <th>State</th>
              <th>Start</th>
              <th>End</th>
              <th><span class="glyphicon glyphicon-time"></span> Duration</th>
            </tr>
          </thead>
          <tbody>
            {{ $p := .process}}
            {{range $index, $operation := .operations}}
            <tr class="{{ $operation.GetResult }}">
              <td class="col-md-1"><a class="btn btn-primary btn-xs" href="/admin/log/process/{{ $p.Action.ID }}/operation/{{ $operation.ID }}"><span class="no-print">See</span> #{{ $operation.ID }}</a></td>
              <td class="col-md-3">{{ $operation.Action.Content }}</td>
              <!-- start state -->
              <td class="col-md-1">
                {{ if $operation.Action.IsRunning }}
                <span class="text-warning">Working</span>
                {{ else }}
                Finished
                {{ end }}
              </td>
              <!-- end state -->
              <td class="col-md-1.5">{{ $operation.GetStartDate }}</td>
              <td class="col-md-1.5">{{ $operation.GetEndDate }}</td>
              <td class="col-md-3">
                {{ if eq $operation.GetDuration "..." }}
                <img src='/static/img/wisply/load.gif' style='height: 20px; width: 20px' />
                {{ else }}
                {{ $operation.GetDuration }}
                {{ end }}
              </td>
            </tr>
            {{end }}
          </tbody>
        </table>
      </div>
    </div>
    {{ end }}
  </div>
</div>
<script>
$(document).ready(function(){
  /**
  * It sends a POST form to an addres
  * @param  {string} address The address of the page
  * @param  {number} ID      The ID of the process
  */
  function sendForm(address, ID) {
    var msg = {
      url: "/admin/harvest/" + address + "/" + ID,
      success: function() {
        window.location="/admin/log";
      },
    };
    wisply.executePostAjax(msg);
  }
  $("#recover-button").click(function(event){
    event.preventDefault();
    console.log($(this).data("id"))
    sendForm("recover", $(this).data("id"));
  });
  $("#finish-button").click(function(event){
    event.preventDefault();
    sendForm("finish", $(this).data("id"));
  });
});
</script>
<script src="/static/js/admin/log/process.js"></script>
