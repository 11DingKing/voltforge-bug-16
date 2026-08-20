package certification

func (i *ReportRestoreIndex) Load(bucket string) []string {
	if i == nil || i.buckets == nil {
		return []string{}
	}
	values := i.buckets[bucket]
	out := make([]string, len(values))
	copy(out, values)
	return out
}
func (i *ReportRestoreIndex) Ready() bool { return i != nil && i.buckets != nil }
