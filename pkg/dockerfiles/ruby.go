package dockerfiles

const Ruby = `FROM ruby:latest

WORKDIR /app

COPY Gemfile Gemfile.lock /app/

RUN bundle install

COPY . /app

RUN groupadd -r rubyapp && useradd -r -g rubyapp rubyapp
USER rubyapp

CMD ["ruby", "app.rb"]`
