describe('Institutions', function() {
  it('goes to the institutions page', function(done) {
      browser
      .url("/institutions")
      .windowHandleMaximize()
      .getTitle(function(err, title) {
          expect(err).toBe(undefined);
          expect(title).toBe('Wisply - Institutions');
      })
      .call(done);
    });
    it("clicks an institution", function(done){
        browser
        .click('a=University of Southampton')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).toBe('University of Southampton');
        })
        .call(done);
    });
    it("sees the institution web page", function(done){
        browser
        .click('a=Web page')
        .getTitle(function(err, title) {
            expect(err).toBe(undefined);
            expect(title).not.toBe('University of Southampton');
        })
        .back()
        .call(done);
    });
});
