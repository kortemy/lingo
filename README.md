lingo
=====

Very basic Golang library for i18n. There are others that do the job, but this is my take on the problem.

Features:
---------
1. Storing messages in JSON files.
2. Support for nested declarations.
2. Detecting language based on Request headers.
3. Very simple to use.

Usage:
------
  1. Create a dir to store translations, and write them in JSON files named [locale].json. For example:
  
      ```    
        en_US.json
        sr_RS.json
        de.json
        ...
      ```
      You can write nested JSON too.
      ```json
        {
          "lingo.example.1" : "Example value 1",
          "lingo.example" : {
            "2" : "Nested example",
            "3" : "Nested example too.",
            "4" : {
                  "inception" : "Double nested?"
            }
          }
        }
      ```
  2. Initialize a Lingo like this:
  
      ```go
        lingo := New("en_US", "path/to/translations/dir")
      ```
  3. Get bundle for specific locale via either `string`: 
  
      ```go
        t := lingo.TranslationsForLocale("sr_RS")
      ```
      This way Lingo will return the bundle for specific locale, or default if given is not found.
      Alternatively (or primarily), you can get it with `*http.Request`:
      
      ```go
        t := lingo.TranslationsForRequest(req)
      ```
      This way Lingo finds best suited locale via `Accept-Language` header, or if there is no match, returns default.
  4. Once you get T instance just fire away!
  
      ```go
        r1 := t1.Value("lingo.example.1") // Example value 1
      	r2 := t1.Value("lingo.example.2") // Nested example
      	r2 := t1.Value("lingo.example.4.inception") // Double nested?
      ```

TODO:
-----
  1. Feature to allow you to get T of specific segment in JSON, and not just string values. That way you can
      group specific messages and extract only that segment where you need it, of different views and/or components.
  2. Write more test cases.
  3. Write better examples that "Example value 1"

Notice:
-------
I am a gopher newb, been hacking around for a couple of months in my spare time, so any comment/contrubution is welcome.
That said, use at your own risk. :D
