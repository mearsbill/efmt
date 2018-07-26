# efmt - this package is a collection of string helper methods

*  efmt.Float2Str(), efmt.Int2Str():  engineering numeric formatting
- Use engineering prefixes for numbers to string conversion
- Supports base 2 mega (2^10 = 1024) conversion.... so 1024 =>1.000K instead of 1.024K
  *
  * efmt.Int2Str()
#
# type Ntag object to manage and print hierchial numbering of data.
#
# eFmt.NewNtag()   // make a new object
func (nt *Ntag) Clone() *Ntag   // clone object to new copy (for a child object)
func (nt *Ntag) SetIndent(indentStr string)  // set the string to use for each indent of hierarch
func (nt *Ntag) Push() // push to deeper level [ 1 = 1.1  1.1=1.1.1 ]
func (nt *Ntag) Pop()  // pop to less deep level [ 1.1.1 => 1.1 ]
func (nt *Ntag) Next() // move to next number at same level [ 1.1.1 => 1.1.2 ]
func (nt *Ntag) String() string        // output the string representation
func (nt *Ntag) Indent() string        // output an "indent string" proportionally to depth of numbering


-
