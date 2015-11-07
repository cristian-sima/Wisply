<div class="page-header">
  <div class="row" >
    <div class="col-lg-12 col-md-12 col-sm-12" >
      <div class="panel panel-default">
        <div class="panel-heading" style="padding-bottom:0px">
          <ul class="breadcrumb">
            <li><a href="/">Home</a></li>
            <li><a href="/about">About</a></li>
            <li class="active">Filters</li>
          </ul>
        </div>
        <section class="panel-body">
          <h1>Filters</h1>

          During the processing of data, Wisply is removing a list of words. These words are removed:
          {{ range $filterName, $filter := .filters }}
          <article>
            <h2>{{ $filterName }}</h2>
            <div class="well">
              {{ range $index, $word := $filter }}
              {{ $word }}, 
              {{ end }}
            </div>
          </ul>
        </article>
        <br />
        {{ end }}
        <article>
          <h2>Other words</h2>
          <ul>
            <li>Empty strings</li>
            <li>Numbers</li>
          </ul>
        </article>
      </section>
    </div>
  </div>
</div>
</div>
