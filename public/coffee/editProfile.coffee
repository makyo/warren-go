# Change the value of the pronouns field to selected pronoun in the chooser.
$('.pronoun-chooser a').click () ->
  $('#pronouns').val $(this).text()
