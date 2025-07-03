FROM rocker/r-ver:4.5.1

RUN apt-get update && apt-get install -y curl gpg
RUN curl -fsSL https://quarto.org/download/latest/quarto-linux-amd64.deb -o quarto.deb && \
    dpkg -i quarto.deb && rm quarto.deb

RUN apt-get update && apt-get install -y \
    libcurl4-openssl-dev \
    libssl-dev \
    libxml2-dev \
    libfontconfig1-dev \
    libfreetype6-dev \
    libharfbuzz-dev \
    libfribidi-dev \
    libpng-dev \
    libtiff5-dev \
    libjpeg-dev \
    pkg-config \
    git \
    pandoc \
    make \
    curl \
    gpg \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /post
COPY . /post

RUN Rscript -e "install.packages('renv'); renv::restore()"

RUN quarto render content.qmd

