### names
Example usage of `names` to randomly assign a name to a [linked list](https://gitlab.com/rshmhrj/data-structs/lists) (similar to Docker & Kubernetes auto naming)
```go
// Initialize an empty linked list
func NewLinkedList(data ... interface{}) *linkedList {
return &linkedList{
name: names.NewGenerator().GenerateShort(),
...
}

l := lists.NewLinkedList("A", "B", "C")
fmt.Printf("%v\n", l)
```

Example Results for `GenerateShort()`:
```bash
the.much.game
as.busy.duty
last.home.blow
ill.fat.web
out.warm.map
not.west.shop
fast.high.till
none.last.roll
only.glad.rush
also.home.lock
true.only.bat
far.big.file
```

Example Results for `Generate()`:
```bash
well.inner.disaster
elsewhere.each.hair
recently.global.farm
nearly.late.chain
essentially.sudden.brain
here.obvious.airport
strange.unable.literature
each.electrical.clue
fresh.inner.housing
often.serious.relative
away.political.confidence
initially.much.classroom
```

### Details

#### NewGenerator() *NameGen
Creates NameGen filled to the brim with nouns, adjectives, and adverbs

#### NameGen.SetDelimiter(delimiter)
Sets the delimiter for output strings

The default delimiter is the `DotDelimiter` resulting in '.' between all words

You can change to the following delimiters:
- DashDelimiter '-'
- BarDelimiter '_'
- CommaDelimiter ','
- SpaceDelimiter ' '


#### NameGen.Generate() string
Randomly chooses three words in the order adverb, adjective, noun with the chosen delimiter separating each word

#### NameGen.GenerateShort() string
Same as Generate except it chooses from words with less than 5 letters

#### NameGen.String() string
Prints the following analysis:
```bash
Nouns: {
 FullLength=1525, ShortLength=438, LongestWord=responsibility, ShortestWord=a 
 Frequency: { 1:1, 2:4, 3:110, 4:323, 5:304, 6:245, 7:189, 8:135, 9:82, 10+:132 }
}
Adjectives: {
 FullLength=527, ShortLength=127, LongestWord=administrative, ShortestWord=ok 
 Frequency: { 1:0, 2:2, 3:29, 4:96, 5:106, 6:88, 7:73, 8:54, 9:38, 10+:41 }
}
Adverbs: {
 FullLength=255, ShortLength=74, LongestWord=significantly, ShortestWord=as 
 Frequency: { 1:0, 2:9, 3:16, 4:49, 5:44, 6:39, 7:28, 8:21, 9:23, 10+:26 }
}
```