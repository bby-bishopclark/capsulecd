from phusion/passenger-full
maintainer Jason Kulatunga <jk17@ualberta.ca>

run apt-get install -y git
run gem install bundler
run npm install -g bower

# copy the application files to the image
workdir /srv/capsulecd
run git clone https://github.com/AnalogJ/capsulecd.git .

run bundle install --with test --without chef ruby python npm
run bundle exec rake spec:javascript
CMD ["bundle", "exec", "capsulecd", "start", "--runner", "circleci", "--source", "github", "--package_type", "javascript"]
