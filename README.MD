# Concurrent duplicate detection with LZW compression
[![Go Report Card](https://goreportcard.com/badge/github.com/shibaeff/CompressionDetection)](https://goreportcard.com/report/github.com/shibaeff/CompressionDetection)

This Golang code was written for my experiment. The aim was to detect duplicate strings/files/sequences ... in the 
concurrent fashion and with the given piece is not a duplicate of previously seen add it to the archive, also
concurrently. To quickly detect the near-duplicates I'm using the [Simhash algo](https://en.wikipedia.org/wiki/SimHash),
compression is done with canonical version of [LZW](https://en.wikipedia.org/wiki/Lempel–Ziv–Welch). 

Sample runs were done on [this](https://ir.shef.ac.uk/cloughie/resources/plagiarism_corpus.html#Citation) dataset - Corpus 
of Plagiarized Short Answers. The work is done by arbitary number of Worker objects. Each of them receives sample strings
via channel, concurrenly computes SimHash and compares it to previosly known instances. If instance is sufficiently unique 
(according to Simhash), the sample is compressed and concatenated to the resulting binary file.
