<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/log">Event log</a></li>
      <li><a href="/admin/log/process/{{ .process.Action.ID }}">Process #{{ .process.Action.ID }}</a></li>
      <li class="active">Operation #{{ .operation.Action.ID }}</li>
    </ul>
  </div>
  <div class="panel-body">
      <span class="text-warning">
				<span class="glyphicon glyphicon-warning-sign"></span>
        This page is not live updated.
			</span>
      <h2>Operation #{{ .operation.Action.ID }}</h2>
      <div class="row">
        <div class="col-lg-4 col-md-4 col-sm-4">
        </div>
        <div class="col-lg-4 col-md-4 col-sm-4">
            <table class="table">
                <tbody>
                    <tr>
                        <td><span class="glyphicon glyphicon-calendar"></span> Start:</td>
                        <td>{{ .operation.GetStartDate }}</td>
                    </tr>
                    <tr>
                        <td><span class="glyphicon glyphicon-calendar"></span> Finish:</td>
                        <td><strong>{{ .operation.GetEndDate }}</strong></td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="col-lg-4 col-md-4 col-sm-4">
            <table class="table">
                <tbody>
                      <td>
                        {{ if .operation.Action.IsRunning }}
                        <span class="label label-warning">Working</span>
                        {{ else }}
                        <span class="label label-success">Finished</span>
                        {{ end }}
                      </td>
                      <td>
                        <strong>
                          {{ if .task }}
                          <a href="/admin/log/process/{{.process.Action.ID}}/operation/{{.operation.ID}}">{{ .operation.Content }}</a>
                          {{ else }}
                          <span class="glyphicon glyphicon-time"></span> {{ .operation.GetDuration }}
                          {{ end }}
                      </strong>
                    </td>
                </tbody>
            </table>
        </div>
      </div>
      <div class="table-responsive">
      <table id="list-tasks" class="table table-hover table-bordered table-condensed">
        <thead>
          <tr>
            <th class="hidden-xs">#</th>
            <th>Content</th>
            <th>State</th>
            <th>Start</th>
            <th>End</th>
            <th>Duration</th>
            <th>Explication</th>
          </tr>
        </thead>
        <tbody>
          {{range $index, $task := .tasks}}
          <tr class="{{ $task.GetResult }}">
            <td class="col-md-1">{{ $task.ID }}</a></td>
            <td class="col-md-1">{{ $task.Action.Content }}</td>
            <!-- start state -->
            <td class="col-md-1">
            {{ if $task.Action.IsRunning }}
            <span class="text-warning">Working</span>
            {{ else }}
            Finished
            {{ end }}
            </td>
            <!-- end state -->
            <td class="col-md-2">{{ $task.GetStartDate }}</td>
            <td class="col-md-2">{{ $task.GetEndDate }}</td>
            <td class="col-md-2">{{ $task.GetDuration }}</td>
            <td class="col-md-2">{{ $task.Explication }}</td>
          </tr>
          {{end }}
        </tbody>
      </table>
    </div>
  </div>
</div>
