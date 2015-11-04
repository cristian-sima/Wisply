<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/institutions">Institutions</a></li>
            <li><a href="/institutions/{{ .institution.ID }}">{{ .institution.Name }}</a></li>
            <li><a href="/institutions/{{ .institution.ID }}/program/{{ .program.GetID }}">{{ .program.GetTitle }}</a></li>
            <li class="active">{{ .module.GetTitle }}</li>
          </ul>
        </div>
        <div class="panel-body">
          <div style="margin:10px">
          <div class="row">
          <h1>{{ .module.GetTitle }}</h1>
          <div class="col-md-9 well">{{ .module.GetContent }}</div>
          <div class="col-md-3">
            <table class="table">
              <tbody>
                <tr>
                  <td>Code</td>
                  <td>{{ .module.GetCode }}</td>
                </tr>
                <tr>
                  <td>Year</td>
                  <td>{{ .module.GetYear }}</td>
                </tr>
                <tr>
                  <td>CATS</td>
                  <td>{{ .module.GetCredits "CATS" }}</td>
                </tr>
                <tr>
                  <td>ECTS</td>
                  <td>{{ .module.GetCredits "ECTS" }}</td>
                </tr>
                <tr>
                  <td>US credits</td>
                  <td>{{ .module.GetCredits "US" }}</td>
                </tr>
              </tbody>
            </table>
          </div>
          </div>
        </div>
          <br />
          <br />
        </div>
      </div>
    </div>
  </div>
</div>
