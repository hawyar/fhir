## FHIR Server




Resource is the basic building block of all exhcangeable FHIR resources. Resource instances are represented as either XML, JSON or RDF and there are currently [145 different resource types](http://hl7.org/fhir/r4/resourcelist.html) defined in the FHIR specification. 



### Operations


 ### Base Resource Content
```json
{
  "resourceType" : "X",
  "id" : "12",
  "meta" : {
    "versionId" : "12",
    "lastUpdated" : "2014-08-18T15:43:30Z",
    "profile" : ["http://example-consortium.org/fhir/profile/patient"],
    "security" : [{
      "system" : "http://terminology.hl7.org/CodeSystem/v3-ActCode",
      "code" : "EMP"
    }],
    "tag" : [{
      "system" : "http://example.com/codes/workflow",
      "code" : "needs-review"
    }]
  },
  "implicitRules" : "http://example-consortium.org/fhir/ehr-plugins",
  "language" : "X"
}
```
 - resourceType: always found in every resource. In XML, this is the name of the root element for the resource
 - id:  defined when the resource is created, and never changed. Only missing when the resource is first created
 - meta.versionId: changes each time any resource contents change (except for the last 3 elements in meta - profile, security and tag)
 - meta.lastUpdated: Changes when the versionId changes. Systems that don't support versions usually don't track lastUpdated either
 - meta.profile: An assertion that the content conforms to a profile. See Extending and Restricting Resources for further discussion. Can be changed as profiles and value sets change or the system rechecks conformance


