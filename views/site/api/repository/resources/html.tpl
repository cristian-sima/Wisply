<table class="table">
  <tbody>
    {{range $index, $record := .records}}
    <tr>
      <td>
          {{ if eq ($record.Keys.Get "relation" | len) 1 }}
            {{range $index, $relation := $record.Keys.Get "relation" }}
              <a class="showTheater"  href="{{ $relation }}">
                <h4>
              {{range $index, $title := $record.Keys.Get "title" }}
              {{ $title }}
              {{ end }}
            </h4>
            {{ end }}
          {{ else }}
          Multirelation
          {{ end }}
          </a>
        {{range $index, $description := $record.Keys.Get "identifier" }}
        <!-- <span class="label label-info">{{ $description }} </span><br /> -->
        {{ end }}

        {{range $index, $description := $record.Keys.Get "description" }}
        {{ $description }}
        {{ end }}
        <div class="formats">
          {{range $format, $number := $record.Keys.ProcessFormats }}
          <span class="label label-default">{{ $format}}{{ if ne $number 1}}<span class=" text-almost-invisible"> <span class="small">&times</span> </span> <span class="">{{ $number }}</span>
          {{ end }}</span> &nbsp;
          {{ end }}
        </div>
        <div class="creators">
          <span class="text-muted">
            {{range $index, $creator := $record.Keys.Get "creator" }}
            {{ $creator }}
            {{ end }}
          </span>
        </div>
      </td>
    </tr>
    {{ end }}
  </tbody>
</table>
