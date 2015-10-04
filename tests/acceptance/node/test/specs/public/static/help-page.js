describe('Help functionality', function() {
  it('loads the help page', function(done) {
      browser
          .windowHandleMaximize()
          .url('/help')
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe('Help');
          })
          .call(done);
  });
  it("cotains the notice about cookies", function(done) {
      browser
      .pause(2000)
      .isVisible(".cc_message").then(function(isVisible) {
        expect(isVisible).toBe(true);
      })
      .call(done);
  });
  it("closes the notification", function(done) {
      browser
      .click(".cc_btn_accept_all")
      .refresh()
      .getCookie("cookieconsent_dismissed").then(function(cookie) {
        expect(cookie).not.toBe(undefined);
        expect(cookie.value).toBe("yes");
      })
      .call(done);
  });
  it("navigates to the legal documents", function (done) {
      browser
      .element('.panel-body')
      .click('a=Legal aspects')
      .pause(1000)
      .isVisible("#legal-aspects").then(function(isVisible) {
        expect(isVisible).toBe(true);
      })
      .call(done);
  });
  it("goes to the cookies", function (done) {
      browser
      .element('ol li h5')
      .click('a=Cookies Policy')
      .getTitle(function(err, title) {
          expect(err).toBe(undefined);
          expect(title).toBe('Wisply Cookies Policy');
      })
      .url(function(err, res) {
              expect(res.value).toBe("http://localhost:8081/cookies");
      })
      .call(done);
  });
  it("jumps to a certain point", function (done) {
      browser
      .element('.panel-body')
      .click('a=What cookies do we use?')
      .pause(1000)
      .isVisible("#what-cookies-do-we-use").then(function(isVisible) {
        expect(isVisible).toBe(true);
      })
      .call(done);
  });
  it("scrolls the page", function (done) {
      browser
      .scroll(0, 500)
      .pause(500)
      .isVisible("#can-I-delete-the-cookies").then(function(isVisible) {
        expect(isVisible).toBe(true);
      })
      .call(done);
  });
  it("goes to top and goes back to legal", function (done) {
      browser
      .scroll(0, 0)
      .pause(500)
      .click('a=Legal aspects')
      .call(done);
  });
  it("it goes to the legal documents", function (done) {
      browser
      .click('a=Terms and conditions')
      .url(function(err, res) {
              expect(res.value).toBe("http://localhost:8081/terms-and-conditions");
      })
      .pause(500)
      .call(done);
  });
  it("it goes to help and home", function (done) {
      browser
      .click('a=Help')
      .pause(1000)
      .click('a=Home')
      .pause(1000)
      .url(function(err, res) {
              expect(res.value).toBe("http://localhost:8081/");
      })
      .call(done);
  });
});
