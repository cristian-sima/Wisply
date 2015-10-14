<div class="panel panel-default">
  <div class="panel-heading" style="padding-bottom:0px">
    <ul class="breadcrumb">
      <li><a href="/admin">Admin</a></li>
      <li><a href="/admin/repositories">Repositories</a></li>
      <li class="active">{{ .repository.Name }}</li>
    </ul></div>
    <div class="panel-body">
      <div class="row" >
        <div class="col-md-8">
          <div class="row">
            <div class="col-md-12">
              <span class="h4">{{ .repository.Name }}</span>
              <span class="repository-status" data-toggle="tooltip" data-placement="right" title="See progress" >{{ .repository.Status }}</span>
              <br />
            </div>
          </div>
          <br /><br /><br />
          <div class="row text-center">
            <div class="col-md-6">
              <div><span class="label label-as-badge big-number label-success">75</span></div>
              <div><h5 class="text-muted">collections</h5></div>
            </div>
            <div class="col-md-6">
              <div><span class="label label-as-badge big-number label-success">4500</span></div>
              <div><h5 class="text-muted">records</h5></div>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div>
            <a href="/repository/{{ .repository.ID }}">Public page</a>
            <br />
            <br />
          </div>
          <div>{{ .repository.Category }} repository.</div>
          <div>Is part of <a href="/admin/institutions/institution/{{ .institution.ID }}">{{ .institution.Name }}</a></div>
        </div>
      </div>
      <br /><br />
      {{ if  or (eq  .repository.Status "unverified") (eq  .repository.Status "verification-failed") }}
      There is no identification available for the repository.
      {{ else }}
      <div class="row" >
        <div class="col-md-6">
          <table class="table">
            <tbody>
              <tr>
                <td>
                  Base URL
                </td>
                <td>
                  <a href="{{ .repository.URL }}">{{ .repository.URL }}</a>
                </td>
              </tr>
              <tr>
                <td>
                  Protocol version
                </td>
                <td>
                  {{ .identification.Protocol }}
                </td>
              </tr>
              <tr>
                <td>
                  Granularity
                </td>
                <td>
                  {{ .identification.Granularity }}
                </td>
              </tr>
              <tr>
                <td>
                  Earliest Datestamp
                </td>
                <td>
                  {{ .identification.EarliestDatestamp }}
                </td>
              </tr>
              <tr>
                <td>
                  Delete Policy
                </td>
                <td>
                  {{ .identification.RecordPolicy }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="col-md-6">
          <p class="text-muted">{{ .repository.Description }}</p>
          <h6>Contact</h6>
          <div class="list-group">
            {{range $index, $email := .identification.AdminEmails}}
            <a href="mailto:{{ $email }}" class="list-group-item">{{ $email }}
            </a>
            {{ end }}
          </div>
        </div>
      </div>
      {{ end }}
      <hr />
      <div>
        <a href="/admin/repositories/repository/{{ .repository.ID }}/advance-options">Advance options</a>
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
    $(".repository-status").each(function(){
      var el = $(this),
      status = wisply.repositoriesModule.GUI.getStatusColor(el.html())
      el.html(status)
    });
    wisply.activateTooltip()
  });
  </script>
