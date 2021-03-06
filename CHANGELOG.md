## Jan 23 2019 - v0.10.1
  - Added MaxIdleConnsPerHost http client option
  - Added TrimSpaces data/udf

## Jan 19 2019 - v0.10.0
  - Added IsNumber helper function
  - Enhance Process Struct to handle unexported fields
    * Added UnexportedFieldHandler hadnler
    * Defined SetUnexportedFieldHandler function
    * Defined IgnoreUnexportedFields default handler
  - Enhanced GetStructMeta to handle unexported fields
    * Added StructMetaFilter
    * Defined DefaultStructMetaFilter default
    * Defined SetStructMetaFilter


## Jan 15 2019 - v0.9.0
   - DownloadWithURL(URL string) (io.ReadCloser, error)
   
## Jan 13 2019 - v0.8.1
  - Moved storage/aws to storage/s3
  - Added lazy s3 bucket creation on upload
  - Added data/udf QueryUnescape 


## Jan 13 2019 - v0.7.0
  - Added ScanStructMethods method
  - Added TryDiscoverValueByKind method
  - Patched AsString UDF
  - Added Base64DecodeText udf
  - Added AnyJSONType for generic interface{} types
  - Added AccountID to cred/config
  - Patched []uint data substitution parsing
    
## Jan 8 2019 - v0.6.5
 - Updated struct to map conversion with honoring tag name

## Jan 7 2019 - v0.6.5
 - Patched non writable map data type mutation

## Jan 6 2019 - v0.6.4
 - Added nested array mutation in data.Map
 - Patched url.resource yaml loading with array structure
 - Patched IndexOf
 - Minor patched
 
## Jan 3 2019 - v0.6.3
 - Added NotFound error with helper functions
 - Updated handling not found on upload logic
 - Added Base64Encode, Base64Decode data udf
 - Added TerminatedSplitN text util function

## Jan 2 2019 - v0.6.2
 - Added FollowRedirects option to http client

## Jan 1 2019 - v0.6.1
 - Patched SortedIterator
 - Patched embedded non pointer struct conversion

## Dec 29 2018 - v0.6.0
 - Added SortedRange, Iterator, SortedIterator to compacted slice
 
## Dec 28 2018 - v0.5.4
 - Added QueryEscape udf
 - Updated handling udf with single quoted literals 

## Dec 27 2018 - v0.5.3
 - Added DecoderFactory method to url.Resource 
 - Patched secret location with URL scheme 

## Dec 26 2018 - v0.5.2
 - Patched KV nested slice conversion 
 - Patched handling unexported fields 
 - Minor patches

## Dec 24 2018 - v0.5.1
 - Patched KV conversion where value was nil
 - Updated secret service location lookup order
 - Minor patches

## Dec 18 2018 - v0.5.0
 - NormalizeKVPairs - to converts slice of KV paris into a map, and map[interface{}]interface{} to map[string]interface{} 
 - Moved stand expandable UDF from neatly project
 - Added data and data/udf documentation

## Dec 7 2018 - v0.4.1
 - Enhanced UDF multi arguments calls
 - Added [] sub map key expression support
 - Patched name with sub references in Map.SetValue

## Dec 7 2018 - v0.4.0
 - Added elapsed/remaining day helper functions: ElapsedDay, ElapsedToday, RemainingToday

## Dec 6 2018 - v0.3.8
 - Patched udf arguments conversion glitch
 - Patched scp service with additional fallback for file scraping
 - Refactor data/map expression parser, added basic arithmetic support
 - Added expansion of struct datatype, patched asEncodableValue
 
 
## Dec 3 2018 - v0.3.1
 - Refactor data/map expression parser, added basic arithmetic support
 - Refactor tokenizer matchers 

## Nov 28 2018 - v0.2.4
 - Patched ToInt, ToFloat conversion with nil pointer

## Nov 24 2018 - v0.2.3
 - Added ToBoolean
 - Streamline ssh Session init
     
## Nov 19 2018 - v0.2.2
 - Added error check for opening shell in ssh Session
 - Enhance SSH termination error

## Nov 8 2018 - v0.2.0

 - Added TimeAt utility method for creating dynamically semantic based time.
 - Added IllegalTokenError, ExpectToken and ExpectTokenOptionallyFollowedBy parsing helpers
 - Added RemainingSequenceMatcher
 - Added SSH stdout buffering with listener frequency flush

## Oct 20 2018 - v0.1.1

 - Added Replace method on data/map.go
 - Added path support to Delete method on on data/map

## Jul 1 2016 (Alpha)

  * Initial Release.
