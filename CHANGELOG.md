WooCommerce for golang Change Log
=================================

## [unreleased]

- New: Batch Product Create/Update/Delete

## 1.0.7

- Bug: update module base to make this code usable 

## 1.0.6 

- FIX: Product Attribute Item was referred to wrong struct

## 1.0.5 

## 1.0.4 under development

-Bug: #1 Fixed signature error if query params include array value

## 1.0.3

- Bug: Fixed date type params validation
- Bug: Fixed Setting group and option properties
- Enh: Add 501 error handling

## 1.0.2

- Bug: Fixed parse string to float64 failed if an empty string
- New: Added data and report service tests
- Bug: Fixed an issue report query parameters did not take effect
- Bug: Fixed report struct error
- Chg: Modify order money field type from string to float64

## 1.0.1

- Enh: Perfect doc
- Enh: Perfect product variation query params validation
- Bug: Fixed All() method isLastPage return error
- Chg: Simplify query params process
- Bug: Fixed Include, Exclude query params type error
- Bug: Fixed shipping zone location endpoint error
- Enh: Set per page is max to 100
- Bug: Fixed is last page check condition
- Enh: Add total and totalPages return in All() method
- Chg: product and product variation price, weight attribute change to float64

## 1.0.0

- Initial release.