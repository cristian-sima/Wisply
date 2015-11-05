<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li class="active">Education</li>
    </ul>
  </div>
  <div class="panel-body">
    <div class="btn-group-sm">
      <a href="/admin/education/subjects/add" class="btn btn-success btn-sm">
        <span class="glyphicon glyphicon-plus"></span> Add subject of study</a>
      </div>
      <br />
      {{ if eq (len .subjects) 0 }}
      There is no subject of study.
      {{ else }}
      <div>
        <div class="row text-center">
          {{range $index, $subject := .subjects }}
          <div class="text-central col-xs-12 col-sm-6 col-md-3 col-ls-2" >
            <a href="/admin/education/subjects/{{ $subject.GetID }}">
              <div style="height:100px;" class="thumbnail">
                <div class="caption">
                  <h3>
                    {{ ($subject.GetName) | html }}
                  </h3>
                </div>
              </div>
            </a>
          </div>
          {{end }}

        </div>
      </div>
      <a class="btn btn-warning" href="/admin/education/analyse">Analyse data</a>
      <br />
      <br />
      {{ end }}
      {{ if eq (.analyses | len) 0 }}
      There are no analyses.
      {{ else }}
      <div class="table-responsive">
        <table id="list-accounts" class="table table-striped table-hover ">
          <thead>
            <tr>
              <th>Start</th>
              <th>End</th>
              <th>Delete</th>
            </tr>
          </thead>
          <tbody>
            {{ $subject := .subject }}
            {{range $index, $analyser := .analyses }}
            <tr>
              <td>{{ $analyser.GetStartDate }}</td>
              <td>{{ $analyser.GetEndDate }}</td>
              <td>
                {{ if $analyser.IsFinished }}
                <a href="#" data-id="{{ $analyser.GetID }}" class="deleteAnalyserButton btn btn-danger btn-xs" >Delete</a>
                {{ else }}
                -
                {{ end }}
              </td>
            </tr>
            {{ end }}
          </tbody>
        </table>
      </div>
      {{ end }}
    </div>
  </div>
  <script src="/static/js/admin/education/home.js"></script>
  <script>
  $(document).ready(function(){
      var module = wisply.getModule("admin-education-home"),
        list = new module.List();
        list.init();
  });
  </script>
