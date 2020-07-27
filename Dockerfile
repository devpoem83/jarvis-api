FROM buildpack-deps:buster-scm

ADD mapper /mapper
ADD profile /profile
ADD public /public
ADD jarvis-api /jarvis-api

CMD /jarvis-api
