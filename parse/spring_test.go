package parse

import (
	"testing"
	"time"

	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
)

var springLog = `
Spring Data Commons Changelog
=============================

Changes in version 2.0.10.RELEASE (2018-09-10)
----------------------------------------------
* DATACMNS-1389 - RepositoryConfigurationExtensionSupport.useRepositoryConfiguration(...) fails in strict configuration mode.
* DATACMNS-1387 - Wrong description for CurrentDateTimeProvider.
* DATACMNS-1386 - SpringDataWebConfiguration cannot be introspected when Jackson is not available.
* DATACMNS-1384 - Add support for java.mysql.Timestamp in AnnotationRevisionMetadata.
* DATACMNS-1383 - Custom extension of Pageable as parameter causes query method to be rejected.
* DATACMNS-1376 - Assure JDK 11 compatibility for DefaultMethodInvokingMethodInterceptor.
* DATACMNS-1375 - In case of failures, AnnotationRepositoryMetadata should be explicit about the offending repository interface.
* DATACMNS-1373 - Align class loading in ClassGeneratingEntityInstantiator with ClassGeneratingPropertyAccessorFactory.
* DATACMNS-1370 - Avoid superflous regex type checks while scanning for custom implementations.
* DATACMNS-1369 - Repository initialization should make sure the aggregate root type gets added to the mapping context.
* DATACMNS-1367 - Add debug logging to better identify repository scanning and initialization.
* DATACMNS-1366 - Investigate performance regressions between 2.0 GA and 2.1 RC2.
* DATACMNS-1364 - BasicPersistentEntity.getPersistentProperty(…) returns the unchecked value in ConcurrentReferenceHashMap.
* DATACMNS-1362 - Broken Links in the Docs 404.
* DATACMNS-1360 - Release 2.0.10 (Kay SR10).
* DATACMNS-1359 - Improve exception message for missing accessors and fields.
* DATACMNS-1351 - Fix typos in reference documentation.
* DATACMNS-1174 - Improve error reporting for not supported repository interfaces.


Changes in version 1.13.15.RELEASE (2018-09-10)
-----------------------------------------------
* DATACMNS-1383 - Custom extension of Pageable as parameter causes query method to be rejected.
* DATACMNS-1370 - Avoid superflous regex type checks while scanning for custom implementations.
* DATACMNS-1367 - Add debug logging to better identify repository scanning and initialization.
* DATACMNS-1362 - Broken Links in the Docs 404.
* DATACMNS-1361 - Release 1.13.15 (Ingalls SR15).
* DATACMNS-1351 - Fix typos in reference documentation.


Changes in version 2.1.0.RC2 (2018-08-20)
-----------------------------------------
* DATACMNS-1377 - ConvertingPropertyAccessor loses its converting power for nested properties.
* DATACMNS-1376 - Assure JDK 11 compatibility for DefaultMethodInvokingMethodInterceptor.
* DATACMNS-1375 - In case of failures, AnnotationRepositoryMetadata should be explicit about the offending repository interface.
* DATACMNS-1373 - Align class loading in ClassGeneratingEntityInstantiator with ClassGeneratingPropertyAccessorFactory.
* DATACMNS-1371 - Avoid excessive component scanning for repository fragments.
* DATACMNS-1370 - Avoid superflous regex type checks while scanning for custom implementations.
* DATACMNS-1369 - Repository initialization should make sure the aggregate root type gets added to the mapping context.
* DATACMNS-1368 - Add support for deferred initialization of repositories.
* DATACMNS-1367 - Add debug logging to better identify repository scanning and initialization.
* DATACMNS-1366 - Investigate performance regressions between 2.0 GA and 2.1 RC2.
* DATACMNS-1364 - BasicPersistentEntity.getPersistentProperty(…) returns the unchecked value in ConcurrentReferenceHashMap.
* DATACMNS-1362 - Broken Links in the Docs 404.
* DATACMNS-1359 - Improve exception message for missing accessors and fields.
* DATACMNS-1358 - Release 2.1 RC2 (Lovelace).
* DATACMNS-1351 - Fix typos in reference documentation.`

func TestSpringChangelog(t *testing.T) {
	logs, err := SpringChangelog("", springLog)
	if assert.NoError(t, err) {
		assert.Len(t, logs, 3)
		assert.Equal(t, model.Version{Major: "2", Minor: "0", Patch: "10", Label: "RELEASE"}, logs[0].Version)
		assert.Equal(t, time.Date(2018, time.September, 10, 0, 0, 0, 0, time.UTC), logs[0].Date)

		if assert.Len(t, logs[0].Issues, 18) {
			assert.Equal(t, model.Info, logs[0].Issues[1].Type)
			assert.Equal(t, "DATACMNS-1387", logs[0].Issues[1].IssueID)
			assert.Equal(t, "Wrong description for CurrentDateTimeProvider.", logs[0].Issues[1].Message)
		}
		assert.Len(t, logs[1].Issues, 6)
		if assert.Len(t, logs[2].Issues, 15) {
			assert.Equal(t, model.Fixed, logs[2].Issues[14].Type)
			assert.Equal(t, "DATACMNS-1351", logs[2].Issues[14].IssueID)
			assert.Equal(t, "Fix typos in reference documentation.", logs[2].Issues[14].Message)
		}
	}
}
