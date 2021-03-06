{{ define "footer" }}
<footer>
  <div class="row">
    <div class="col-lg-12">
      <ul class="list-unstyled">
        <li>
          <a href="https://wisplyblog.wordpress.com/" target="_blank" >Blog</a>
        </li>
        <li>
          <a href="https://www.facebook.com/wisply" target="_blank">Facebook</a>
        </li>
        <li>
          <a href="https://twitter.com/wisplyOfficial" target="_blank">Twitter</a>
        </li>
        <li>
          <a href="/developer">Developers &amp; Research</a>
        </li>
        <li>
          <a href="/help">Help</a>
        </li>
        <li>
          <div class="dropup">
            <span role="button" class="hover link dropdown-toggle" id="moreOptionsFooter" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
              <a class="hover">More<span class="caret"></span></a>
            </span>
            <ul class="dropdown-menu" aria-labelledby="moreOptionsFooter">
              <li><a href="/thank-you">Thank you</a></li>
              <li><a href="/accessibility">Accessibility</a></li>
              <li><a href="/privacy">Privacy Policy</a></li>
              <li><a href="/contact">Contact</a></li>
            </ul>
          </div>
        </li>
        <li class="pull-right">
          <a href="javascript:scroll(0,0)">Back to top</a>
        </li>
      </ul>
      {{ if .indicateLastModification }}
      <p>
        This page has been modified on {{ .lastModification }}.
      </p>
      {{ end }}
      <!--<p>We have done the best to express our <a href="/licence" rel="nofollow">Licence</a>.</p> -->
    </div>
  </div>
</footer>
{{ end }}
