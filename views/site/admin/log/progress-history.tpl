<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/log">Event log</a></li>
      <li><a href="/admin/log/process/{{ .process.Action.ID }}">Process #{{ .process.Action.ID }}</a></li>
      <li class="active">History</li>
    </ul>
  </div>
  <div class="panel-body">
      <div>
          <span class="text-warning warning-notice">
  				<span class="glyphicon glyphicon-warning-sign"></span>
          This page is not live updated.
  			</span>
      </div>
      <h2>Process #{{ .process.Action.ID }}</h2>
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
                        <td><strong><a href="/admin/repositories/repository/{{ .harvestProcess.GetRepository.ID }}">{{ .harvestProcess.GetRepository.Name }}</strong></a></td>
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
                          {{ if .process.Action.IsRunning }}
                          <span class="label label-warning">Working</span>
                          {{ else }}
                          <span class="label label-success">Finished</span>
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
    <div class="no-print">
      <a href="/admin/log/process/{{ .process.Action.ID }}">See overview</a> <br />
      <a href="/admin/log/process/{{ .process.Action.ID }}/advance-options">Advance options</a>
    </div>
    <br />
    <hr/>
    <br />
    <div class="well tex-left" id="history" >
    {{ if .process.Action.IsRunning }}
        <span class="glyphicon glyphicon-calendar"></span> <strong>Now</strong>: The process is <span class="label label-warning">working</span> ...
    {{ else }}
        <span class="glyphicon glyphicon-calendar"></span>  {{ .process.GetEndDate }}: The process has <span class="label label-success">finished</span>
    {{ end }}
    </div>
    {{ $p := .process}}
    {{range $index, $operation := .operations}}
    {{ $tasks := $operation.GetTasks }}
    <div class="well text-left">
    {{ if $operation.Action.IsRunning }}
        <span class="glyphicon glyphicon-calendar"></span> <strong>Now</strong>: The operation <a href="/admin/log/process/{{ $p.Action.ID }}/operation/{{ $operation.ID }}">{{ $operation.Action.Content }}</a> is <span class="label label-warning">working</span> ...
    {{ else }}
        <span class="glyphicon glyphicon-calendar"></span>  {{ $operation.GetEndDate }}: The operation <a href="/admin/log/process/{{ $p.Action.ID }}/operation/{{ $operation.ID }}">{{ $operation.Action.Content }}</a> has <span class="label label-success">finished</span> in
            <span class="glyphicon glyphicon-time"></span> {{ $operation.GetDuration }}.
    {{ end }}
    </div>
    <div class="print-div">
    <div class="panel panel-{{ $operation.GetResult }}">
        <div class="panel-heading"><a href="/admin/log/process/{{ $p.Action.ID }}/operation/{{ $operation.ID }}">#{{ $operation.ID }} </a> Operation {{ $operation.Action.Content }}
        </div>
        <div class="panel-body">
          <br />
          <div class="table-responsive">
            <table class="list-tasks table table-bordered table-condensed">
              <thead>
                <tr>
                  <th># Task</th>
                  <th>Content</th>
                  <th>State</th>
                  <th>Start</th>
                  <th>End</th>
                  <th>Duration</th>
                  <th>Explication</th>
                </tr>
              </thead>
              <tbody>
                {{range $index, $task := $tasks}}
                <tr class="{{ $task.GetResult }}">
                  <td class="col-md-1">{{ $task.ID }}</a></td>
                  <td class="col-md-2">{{ $task.Action.Content }}</td>
                  <!-- start state -->
                  <td class="col-md-1">
                    {{ if $task.Action.IsRunning }}
                    <span class="text-warning">Working</span>
                    {{ else }}
                    Finished
                    {{ end }}
                  </td>
                  <!-- end state -->
                  <td class="col-md-1.5">{{ $task.GetStartDate }}</td>
                  <td class="col-md-1.5">{{ $task.GetEndDate }}</td>
                  <td class="col-md-2">
                    {{ if eq $task.GetDuration "..." }}
                    <img src='/static/img/wisply/load.gif' style='height: 20px; width: 20px' />
                    {{ else }}
                    {{ $task.GetDuration }}
                    {{ end }}
                  </td>
                  <td class="col-md-2">
                    {{ $len := $task.Explication | len }}
                      {{ if gt $len 70 }}
                      <a data-explication="{{ $task.Explication | html }}" class="see-full-explication" href="#">See full explication</a>
                    {{ else }}
                      {{ $task.Explication }}
                    {{ end }}
                  </td>
                </tr>
                {{end }}
              </tbody>
            </table>
          </div>
          </div>
        </div>
      </div>
      <div class="well text-left">
        <strong><span class="glyphicon glyphicon-calendar"></span> {{ $operation.GetStartDate }}</strong>: The  operation <a href="/admin/log/process/{{ $p.Action.ID }}/operation/{{$operation.ID }}">{{$operation.Action.Content}}</a> has <span class="label label-default">started</span>
      </div>
      {{ end }}
    <div class="well text-left">
      <strong><span class="glyphicon glyphicon-calendar"></span> {{ .process.GetStartDate }}</strong>: The process has <span class="label label-default">started</span>
    </div>
  </div>
</div>
<script src="/static/js/admin/log/process.js"></script>
<script src="/static/js/admin/log/operation.js"></script>
