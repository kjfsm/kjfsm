FROM ubuntu
RUN mkdir /myvol
RUN echo "hello volume" > /myvol/greeting
VOLUME [ "/myvol" ]
RUN echo "test"
RUN echo /myvol/greeting
