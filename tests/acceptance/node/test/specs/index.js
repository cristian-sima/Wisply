describe('Testing static pages', function() {

  it('loads the home page', function(done) {
      browser
          .url('/')
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe('Wisply - Building the hive of education');
          })
          .call(done);
  });


  it('loads the about page', function(done) {
      browser
          .url('/about')
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe('About Wisply');
          })
          .call(done);
  });

  it('loads the webscience page', function(done) {
      browser
          .url('/webscience')
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe('Webscience');
          })
          .call(done);
  });

  it('loads the contact page', function(done) {
      browser
          .url('/contact')
          .getTitle(function(err, title) {
              expect(err).toBe(undefined);
              expect(title).toBe('Contact Wisply');
          })
          .call(done);
  });

  // Help
  describe('Auth system', function() {
    it('loads the help page', function(done) {
        browser
            .url('/help')
            .getTitle(function(err, title) {
                expect(err).toBe(undefined);
                expect(title).toBe('Help');
            })
            .call(done);
    });
    // Legal documents
    describe('Legal documents', function() {
      it('loads the privacy policy page', function(done) {
          browser
              .url('/privacy')
              .getTitle(function(err, title) {
                  expect(err).toBe(undefined);
                  expect(title).toBe('Wisply Privacy Policy');
              })
              .call(done);
      });
    });
  });
  // Auth
  describe('Auth system', function() {
    it('loads the login page', function(done) {
        browser
            .url('/auth/login')
            .getTitle(function(err, title) {
                expect(err).toBe(undefined);
                expect(title).toBe('Login to Wisply');
            })
            .call(done);
    });
    it('loads the register page', function(done) {
        browser
            .url('/auth/register')
            .getTitle(function(err, title) {
                expect(err).toBe(undefined);
                expect(title).toBe('Create a new account');
            })
            .call(done);
    });
  });

});
