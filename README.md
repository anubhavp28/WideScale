# WideScale
WideScale is a full-text indexing and searching engine, written in golang. WideScale is solely for educational purposes.
It provides a simple API to search for words or group of words, inside large quantity of text spread across multiple documents. 
Internally, WideScale uses a <a href="https://en.wikipedia.org/wiki/Inverted_index">Inverted Index</a>, similar to 
<a href="https://en.wikipedia.org/wiki/Elasticsearch"> ElasticSearch</a>. 
For more information, see <a href="https://www.elastic.co/guide/en/elasticsearch/guide/current/inverted-index.html">this article.</a>

**Let me know if you guys have any suggestions.**
### Why use a special data structure (Inverted Index) ?
I found Inverted Index while I was reading about ElasticSearch. To understand why use it, here is excerpt from wikipedia article about it - 
>> When dealing with a small number of documents, it is possible for the
 full-text-search engine to directly scan the contents of the documents 
with each query, a strategy called "serial scanning". This is what some tools, such as grep, do when searching.

>>However, when the number of documents to search is potentially 
large, or the quantity of search queries to perform is substantial, the 
problem of full-text search is often divided into two tasks: indexing 
and searching. The indexing stage will scan the text of all the 
documents and build a list of search terms (often called an index).
 In the search stage, when performing a specific query, only the index 
is referenced, rather than the text of the original documents.

I really didn't think I could do a better explanation than that. 

## Installation
* Install golang ([Instructions](https://golang.org/doc/install)). Add `go` installation path to your PATH environment variable.
* Download [mux](https://github.com/gorilla/mux).
  ```
  > go get github.com/gorilla/mux
  ```
* Downlaod widescale
  ```
  > go get github.com/anubhavp28/WideScale/
  ```
* Install widescale  
  ```
  > go install github.com/anubhavp28/WideScale/
  ```

## Usage
* To start the server, simply run:

  ```
  > cd $(go env GOPATH)/bin
  > widescale <path-to-dir-containing-txt-files-to-index>
  ```

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.
  
  
